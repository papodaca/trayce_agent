package internal

import "C"
import (
	"fmt"
	"os"
	"runtime"

	"github.com/aquasecurity/libbpfgo"
	"github.com/evanrolfe/dockerdog/internal/models"
)

const (
	bufPollRateMs = 200
	sslLibPath    = "/usr/lib/x86_64-linux-gnu/libssl.so.3"
	libcLibPath   = "/usr/lib/x86_64-linux-gnu/libc.so.6"
)

type BPFAgent struct {
	bpfProg           *BPFProgram
	sockets           models.SocketMap
	interuptChan      chan int
	dataEventsChan    chan []byte
	connectEventsChan chan []byte
	debugEventsChan   chan []byte
	dataEventsBuf     *libbpfgo.RingBuffer
	connectEventsBuf  *libbpfgo.RingBuffer
	debugEventsBuf    *libbpfgo.RingBuffer
}

func NewBPFAgent(bpfBytes []byte, btfFilePath string, libSslPath string) *BPFAgent {
	bpfProg, err := NewBPFProgramFromBytes(bpfBytes, btfFilePath, "")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	// probe_entry_SSL_read
	// Entry gives: HTTP/1.1 301 Moved Permanently..
	bpfProg.AttachToUProbe("probe_entry_SSL_read", "SSL_read", libSslPath)
	bpfProg.AttachToURetProbe("probe_ret_SSL_read", "SSL_read", libSslPath)

	bpfProg.AttachToUProbe("probe_entry_SSL_read_ex", "SSL_read_ex", libSslPath)
	bpfProg.AttachToURetProbe("probe_ret_SSL_read_ex", "SSL_read_ex", libSslPath)

	// probe_entry_SSL_write
	// Return gives: GET / HTTP/1.1..
	bpfProg.AttachToUProbe("probe_entry_SSL_write", "SSL_write", libSslPath)
	bpfProg.AttachToURetProbe("probe_ret_SSL_write", "SSL_write", libSslPath)

	// kprobe connect
	funcName := fmt.Sprintf("__%s_sys_connect", ksymArch())
	bpfProg.AttachToKProbe("probe_connect", funcName)
	bpfProg.AttachToKRetProbe("probe_ret_connect", funcName)

	return &BPFAgent{
		bpfProg:           bpfProg,
		sockets:           models.NewSocketMap(),
		interuptChan:      make(chan int),
		dataEventsChan:    make(chan []byte),
		connectEventsChan: make(chan []byte),
		debugEventsChan:   make(chan []byte),
	}
}

func (agent *BPFAgent) ListenForEvents(outputChan chan models.MsgEvent) {
	// DataEvents ring buffer
	var err error
	agent.dataEventsBuf, err = agent.bpfProg.BpfModule.InitRingBuf("data_events", agent.dataEventsChan)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	// SocketAddrEvents ring buffer
	agent.connectEventsBuf, err = agent.bpfProg.BpfModule.InitRingBuf("connect_events", agent.connectEventsChan)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	// DebugEvents ring buffer
	agent.debugEventsBuf, err = agent.bpfProg.BpfModule.InitRingBuf("debug_events", agent.debugEventsChan)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	agent.dataEventsBuf.Poll(bufPollRateMs)
	agent.connectEventsBuf.Poll(bufPollRateMs)
	agent.debugEventsBuf.Poll(bufPollRateMs)

	for {
		// Check if the interrupt signal has been received
		select {
		case <-agent.interuptChan:
			return

		case payload := <-agent.dataEventsChan:
			event := models.DataEvent{}
			event.Decode(payload)
			// fmt.Println("[DataEvent] Received ", event.DataLen, "bytes, type:", event.Type(), ", PID:", event.Pid, ", TID:", event.Tid, "FD: ", event.Fd)
			// fmt.Println(hex.Dump(eventPayload))

			eventPayload := event.Payload()
			if len(eventPayload) > 256 {
				eventPayload = eventPayload[0:128]
			}

			// Fetch its socket
			socket, exists := agent.sockets[event.Key()]
			if !exists {
				continue
			}
			outputChan <- models.NewMsgEvent(&event, socket)

		case payload := <-agent.connectEventsChan:
			event := models.ConnectEvent{}
			event.Decode(payload)
			// fmt.Println("[ConnectEvent] Received ", len(payload), "bytes", "PID:", event.Pid, ", TID:", event.Tid, "FD: ", event.Fd, ", ", event.IPAddr(), ":", event.Port, " local? ", event.Local)

			// Save the event to the map
			agent.sockets.ParseConnectEvent(&event)
		case _ = <-agent.debugEventsChan:
			continue
			// fmt.Println("[DebugEvent] Received", len(payload), "bytes")
			// fmt.Println(hex.Dump(payload))
		}
	}
}

func (agent *BPFAgent) Close() {
	// agent.dataEventsBuf.Stop()
	// agent.dataEventsBuf.Close()
	// agent.debugEventsBuf.Stop()
	// agent.debugEventsBuf.Close()
	// agent.socketAddrEventsBuf.Stop()
	// agent.socketAddrEventsBuf.Close()

	agent.interuptChan <- 1
	// close(agent.interuptChan)
	// close(agent.socketAddrEventsChan)
	// close(agent.dataEventsChan)

	agent.bpfProg.Close()
}

func ksymArch() string {
	switch runtime.GOARCH {
	case "amd64":
		return "x64"
	case "arm64":
		return "arm64"
	default:
		panic("unsupported architecture")
	}
}
