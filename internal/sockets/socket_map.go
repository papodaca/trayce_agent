package sockets

import (
	"fmt"
	"sync"

	"github.com/evanrolfe/dockerdog/internal/bpf_events"
)

// SocketMap tracks sockets which have been observed in ebpf
type SocketMap struct {
	mu            sync.Mutex
	sockets       map[string]SocketI
	flowCallbacks []func(Flow)
	socketKeys    []string
}

func NewSocketMap() *SocketMap {
	m := SocketMap{
		sockets:    make(map[string]SocketI),
		socketKeys: []string{},
	}
	return &m
}

func (m *SocketMap) GetSocket(key string) (SocketI, bool) {
	socket, exists := m.sockets[key]
	return socket, exists
}

func (m *SocketMap) AddSocket(socket SocketI) {
	socket.AddFlowCallback(func(flow Flow) {
		for _, callback := range m.flowCallbacks {
			callback(flow)
		}
	})

	m.sockets[socket.Key()] = socket
}

func (m *SocketMap) AddFlowCallback(callback func(Flow)) {
	m.flowCallbacks = append(m.flowCallbacks, callback)
}

func (m *SocketMap) ProcessConnectEvent(event bpf_events.ConnectEvent) {
	m.mu.Lock()
	defer m.mu.Unlock()

	socket, exists := m.GetSocket(event.Key())

	if !exists {
		fmt.Println("[SocketMap] Connect - creating socket for:", event.Key())
		// TODO: This should first create an SocketUnknown, then change it to SocketHttp11 once we can detect the protocol
		socket := NewSocketHttp11(&event)
		m.AddSocket(&socket)
	} else {
		fmt.Println("[SocketMap] Connect - found socket for:", event.Key())
		socket.ProcessConnectEvent(&event)
	}
}

func (m *SocketMap) ProcessDataEvent(event bpf_events.DataEvent) {
	m.mu.Lock()
	defer m.mu.Unlock()

	socket, exists := m.GetSocket(event.Key())

	if !exists {
		panic("[SocketMap] not socket found for event") // This should never happen
	}

	fmt.Println("[SocketMap] DataEvent - found socket for:", event.Key(), "/", event.Rand)
	socket.ProcessDataEvent(&event)
}

func (m *SocketMap) ProcessCloseEvent(event bpf_events.CloseEvent) {
	m.mu.Lock()
	defer m.mu.Unlock()
	fmt.Println("[SocketMap] CloseEvent - deleting socket for:", event.Key())
	delete(m.sockets, event.Key())
}

func (m *SocketMap) Debug() {
	socketLine := ""
	for _, socket := range m.sockets {
		socketLine += socket.Key() + ", "
	}
	fmt.Println(socketLine)
}
