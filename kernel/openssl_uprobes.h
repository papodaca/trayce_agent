// go:build exclude

// Key is thread ID (from bpf_get_current_pid_tgid).
// Value is a pointer to the data buffer argument to SSL_write/SSL_read.
struct {
  __uint(type, BPF_MAP_TYPE_HASH);
  __type(key, u64);
  __type(value, struct active_buf);
  __uint(max_entries, 1024);
} active_ssl_read_args_map SEC(".maps");

struct {
  __uint(type, BPF_MAP_TYPE_HASH);
  __type(key, u64);
  __type(value, struct active_buf);
  __uint(max_entries, 1024);
} active_ssl_write_args_map SEC(".maps");

int process_ssl_read_entry(struct pt_regs *ctx, bool is_ex_call) {
    u64 current_pid_tgid = bpf_get_current_pid_tgid();
    u32 pid = current_pid_tgid >> 32;
    bpf_printk("SSL_read ENTRY current_pid_tgid: %d", current_pid_tgid);
    // Check if PID is intercepted
    u32 *local_ip = bpf_map_lookup_elem(&intercepted_pids, &pid);
    if (local_ip == NULL) {
        return 0;
    }

    void *ssl = (void *)PT_REGS_PARM1(ctx);
    // https://github.com/openssl/openssl/blob/OpenSSL_1_1_1-stable/crypto/bio/bio_local.h
    struct ssl_st ssl_info;
    int res = bpf_probe_read_user(&ssl_info, sizeof(ssl_info), ssl);
    if (res < 0) {
        bpf_printk("SSL_read enty bpf_probe_read_user ssl_info failed: %d", res);
    }

    u32 fd = get_fd_from_libssl_read(ssl_info, pid, current_pid_tgid);

    // Save the FD incase SSL_Read or SSL_Write need it
    u64 ssl_ptr = (u64) ssl;
    bpf_map_update_elem(&ssl_fd_map, &ssl_ptr, &fd, BPF_ANY);

    // Workaround: for incoming SSL requests to Ruby servers, for some reason bio_r.num is always = -1
    // so we dont have access to the FD number.
    if (fd == -1) {
        fd = 0;
    }

    const char *buf = (const char *)PT_REGS_PARM2(ctx);
    struct active_buf active_buf_t;
    __builtin_memset(&active_buf_t, 0, sizeof(active_buf_t));
    active_buf_t.fd = fd;
    active_buf_t.version = ssl_info.version;
    active_buf_t.buf = buf;
    active_buf_t.ssl_info = (const struct ssl_st *)PT_REGS_PARM1(ctx);
    active_buf_t.ssl_ptr = (u64)ssl;

    if (is_ex_call) {
        size_t *ssl_ex_len_ptr = (size_t *)PT_REGS_PARM4(ctx);
        active_buf_t.ssl_ex_len_ptr = ssl_ex_len_ptr;
    }

    bpf_map_update_elem(&active_ssl_read_args_map, &current_pid_tgid, &active_buf_t, BPF_ANY);
    return 0;
}

int process_ssl_read_return(struct pt_regs *ctx, bool is_ex_call) {
  u64 current_pid_tgid = bpf_get_current_pid_tgid();
  u32 pid = current_pid_tgid >> 32;
  // bpf_printk("openssl uretprobe/SSL_read pid :%d\n", pid);

  struct active_buf *active_buf_t = bpf_map_lookup_elem(&active_ssl_read_args_map, &current_pid_tgid);

  if (active_buf_t != NULL) {
    const char *buf;
    u32 fd = active_buf_t->fd;

    struct ssl_st ssl_info;
    int res2 = bpf_probe_read_user(&ssl_info, sizeof(ssl_info), active_buf_t->ssl_info);
    if (res2 < 0) {
        bpf_printk("SSL_read RETURN bpf_probe_read_user ssl_info failed: %d", res2);
    }

    u32 fd2 = get_fd_from_libssl_read(ssl_info, pid, current_pid_tgid);

    int res = (int)PT_REGS_RC(ctx);

    size_t buf_len;
    if (is_ex_call) {
        bpf_probe_read(&buf_len, sizeof(buf_len), active_buf_t->ssl_ex_len_ptr);
    } else {
        buf_len = (size_t)PT_REGS_RC(ctx);
    }

    if (res <= 0) {
        return 0;
    }

    s32 version = active_buf_t->version;
    bpf_probe_read(&buf, sizeof(const char *), &active_buf_t->buf);

    process_data(ctx, current_pid_tgid, kSSLRead, buf, buf_len, fd);
  }
  bpf_map_delete_elem(&active_ssl_read_args_map, &current_pid_tgid);

  return 0;
}

int process_ssl_write_entry(struct pt_regs *ctx, bool is_ex_call) {
    u64 current_pid_tgid = bpf_get_current_pid_tgid();
    u32 pid = current_pid_tgid >> 32;
    bpf_printk("SSL_write ENTRY current_pid_tgid: %d", current_pid_tgid);
    // Check if PID is intercepted
    u32 *pid_intercepted = bpf_map_lookup_elem(&intercepted_pids, &pid);
    if (pid_intercepted == NULL) {
    return 0;
    }

    void *ssl = (void *)PT_REGS_PARM1(ctx);

    // https://github.com/openssl/openssl/blob/OpenSSL_1_1_1-stable/crypto/bio/bio_local.h
    struct ssl_st ssl_info;
    bpf_probe_read_user(&ssl_info, sizeof(ssl_info), ssl);

    u32 fd = get_fd_from_libssl_write(ssl_info, pid, (u64) ssl);

    // Workaround: see comment in process_ssl_read_entry()
    if (fd == -1) {
        fd = 0;
    }

    const char *buf = (const char *)PT_REGS_PARM2(ctx);
    struct active_buf active_buf_t;
    __builtin_memset(&active_buf_t, 0, sizeof(active_buf_t));
    active_buf_t.fd = fd;
    active_buf_t.version = ssl_info.version;
    active_buf_t.buf = buf;
    active_buf_t.ssl_ptr = (u64)ssl;

    if (is_ex_call) {
        size_t *ssl_ex_len_ptr = (size_t *)PT_REGS_PARM4(ctx);
        size_t ssl_ex_len;
        bpf_probe_read(&ssl_ex_len, sizeof(ssl_ex_len), ssl_ex_len_ptr);
        bpf_printk("SSL_write_ex FD:%d, ex_len: %d\n", fd, ssl_ex_len);
        active_buf_t.ssl_ex_len_ptr = ssl_ex_len_ptr;
    }

  bpf_map_update_elem(&active_ssl_write_args_map, &current_pid_tgid, &active_buf_t, BPF_ANY);

  return 0;
}

int process_ssl_write_return(struct pt_regs *ctx, bool is_ex_call) {
  u64 current_pid_tgid = bpf_get_current_pid_tgid();
  u32 pid = current_pid_tgid >> 32;

  // Send entry data from map
  // bpf_printk("openssl uretprobe/SSL_write_ex pid :%d\n", pid);
  struct active_buf *active_buf_t = bpf_map_lookup_elem(&active_ssl_write_args_map, &current_pid_tgid);

  if (active_buf_t != NULL) {
    const char *buf;
    u32 fd = active_buf_t->fd;
    s32 version = active_buf_t->version;
    bpf_probe_read(&buf, sizeof(const char *), &active_buf_t->buf);
    // bpf_printk("SSL_writ 1--------------------> ssl: %d, fd: %d", ssl);
    size_t buf_len;
    if (is_ex_call) {
      bpf_probe_read(&buf_len, sizeof(buf_len), active_buf_t->ssl_ex_len_ptr);
    } else {
      buf_len = (size_t)PT_REGS_RC(ctx);
    }

    process_data(ctx, current_pid_tgid, kSSLWrite, buf, buf_len, fd);
  }
  bpf_map_delete_elem(&active_ssl_write_args_map, &current_pid_tgid);
  return 0;
}
/***********************************************************
 * BPF uprobes
 ***********************************************************/
// int SSL_read(SSL *s, void *buf, int num)
SEC("uprobe/SSL_read")
int probe_entry_SSL_read(struct pt_regs *ctx) {
  return process_ssl_read_entry(ctx, false);
}

SEC("uretprobe/SSL_read")
int probe_ret_SSL_read(struct pt_regs *ctx) {
  return process_ssl_read_return(ctx, false);
}

// int SSL_read_ex(SSL *ssl, void *buf, size_t num, size_t *readbytes);
SEC("uprobe/SSL_read_ex")
int probe_entry_SSL_read_ex(struct pt_regs *ctx) {
  return process_ssl_read_entry(ctx, true);
}

SEC("uretprobe/SSL_read_ex")
// int SSL_read_ex(SSL *ssl, void *buf, size_t num, size_t *readbytes);
int probe_ret_SSL_read_ex(struct pt_regs *ctx) {
  return process_ssl_read_return(ctx, true);
}

// int SSL_write(SSL *ssl, const void *buf, int num);
SEC("uprobe/SSL_write")
int probe_entry_SSL_write(struct pt_regs *ctx) {
  return process_ssl_write_entry(ctx, false);
}

SEC("uretprobe/SSL_write")
int probe_ret_SSL_write(struct pt_regs *ctx) {
  return process_ssl_write_return(ctx, false);
}

SEC("uprobe/SSL_write_ex")
int probe_entry_SSL_write_ex(struct pt_regs *ctx) {
  return process_ssl_write_entry(ctx, true);
}

SEC("uretprobe/SSL_write_ex")
int probe_ret_SSL_write_ex(struct pt_regs *ctx) {
  return process_ssl_write_return(ctx, true);
}
