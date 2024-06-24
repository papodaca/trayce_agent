package sockets

import (
	"encoding/hex"
	"fmt"
	"sync"

	"github.com/evanrolfe/trayce_agent/internal/bpf_events"
)

// SocketMap tracks sockets which have been observed in ebpf
type SocketMap struct {
	mu            sync.Mutex
	sockets       map[string]SocketI
	flowCallbacks []func(Flow)
}

func NewSocketMap() *SocketMap {
	m := SocketMap{
		sockets: make(map[string]SocketI),
	}
	return &m
}

func (m *SocketMap) AddFlowCallback(callback func(Flow)) {
	m.flowCallbacks = append(m.flowCallbacks, callback)
}

func (m *SocketMap) ProcessConnectEvent(event bpf_events.ConnectEvent) {
	m.mu.Lock()
	defer m.mu.Unlock()

	socket, exists := m.getSocket(event.Key())

	if !exists {
		fmt.Println("[SocketMap] Connect - creating socket for:", event.Key())
		// TODO: This should first create an SocketUnknown, then change it to SocketHttp11 once we can detect the protocol
		socket := NewSocketUnknown(&event)
		m.setSocket(&socket)
	} else {
		fmt.Println("[SocketMap] Connect - found socket for:", event.Key())
		socket.ProcessConnectEvent(&event)
	}
}

func (m *SocketMap) ProcessDataEvent(event bpf_events.DataEvent) {
	m.mu.Lock()
	defer m.mu.Unlock()

	var socket SocketI

	// For some reason HTTPS requests to ruby never have the FD value set, so instead we use the ssl_ptr value as the fd
	// so that we can at least correleate requests with responses, but we wont have a connect event so the src & dst will be 0.0.0.0
	if event.Fd == 0 && event.SslPtr > 0 {
		event.Fd = uint32(event.SslPtr)
	}
	green := "\033[92m"
	reset := "\033[0m"

	fmt.Println(string(green), "[DataEvent]", string(reset), " Received ", event.DataLen, "bytes, source:", event.Source(), ", PID:", event.Pid, ", TID:", event.Tid, "FD: ", event.Fd, " ssl_ptr:", event.SslPtr)
	fmt.Print(hex.Dump(event.PayloadTrimmed(256)))

	socket = m.getOrCreateSocket(event)

	// If the socket is unknown, try to detect the protocol, if not detection then drop
	// but if detected then convert it to the protocol socket
	unkownSocket, isUnknown := socket.(*SocketUnknown)
	if isUnknown {
		protocol := detectProtocol(event.Payload())
		fmt.Println("[SocketMap] detected protocol:", protocol)
		switch protocol {
		case Unknown:
			return
		case HTTP:
			newSocket := NewSocketHttp11FromUnknown(unkownSocket)
			m.setSocket(&newSocket)
			socket = &newSocket
		case HTTP2:
			newSocket := NewSocketHttp2FromUnknown(unkownSocket)
			m.setSocket(&newSocket)
			socket = &newSocket
		}
	}

	socket.ProcessDataEvent(&event)
}

func (m *SocketMap) ProcessCloseEvent(event bpf_events.CloseEvent) {
	m.mu.Lock()
	defer m.mu.Unlock()
	_, exists := m.getSocket(event.Key())
	if exists {
		fmt.Println("[SocketMap] CloseEvent - deleting socket for:", event.Key())
		delete(m.sockets, event.Key())
	}
}

func (m *SocketMap) Debug() {
	socketLine := ""
	for _, socket := range m.sockets {
		socketLine += socket.Key() + ", "
	}
	fmt.Println(socketLine)
}

func (m *SocketMap) getSocket(key string) (SocketI, bool) {
	socket, exists := m.sockets[key]
	return socket, exists
}

func (m *SocketMap) getOrCreateSocket(event bpf_events.DataEvent) SocketI {
	socket, exists := m.getSocket(event.Key())
	if exists {
		fmt.Println("Found socket:", socket.Key())
		return socket
	}

	// If we can't find a socket then lets just create one with dst IP set to 0.0.0.0 becuause thats better than nothing
	newSocket := NewSocketUnknown(&bpf_events.ConnectEvent{
		Pid:  event.Pid,
		Tid:  event.Tid,
		Fd:   event.Fd,
		Ip:   0,
		Port: 80,
	})
	return &newSocket
}

func (m *SocketMap) setSocket(socket SocketI) {
	socket.AddFlowCallback(func(flow Flow) {
		for _, callback := range m.flowCallbacks {
			callback(flow)
		}
	})

	m.sockets[socket.Key()] = socket
}
