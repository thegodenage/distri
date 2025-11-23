package handler

import "sync"

type (
	Manager struct {
		remoteHandlers []RemoteHandler
		mu             sync.Mutex
	}
)

func NewManager() Manager {
	return Manager{}
}

func (m *Manager) RegisterRemoteHandler(han RemoteHandler) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.remoteHandlers = append(m.remoteHandlers, han)
}
