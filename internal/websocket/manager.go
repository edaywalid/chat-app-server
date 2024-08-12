package websocket

import (
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	UserID uuid.UUID
	Conn   *websocket.Conn
}

type Manager struct {
	clients    map[uuid.UUID]*Client
	clientsMux sync.RWMutex
}

func NewManger() *Manager {
	return &Manager{
		clients: make(map[uuid.UUID]*Client),
	}
}

func (m *Manager) AddClient(UserID uuid.UUID, conn *websocket.Conn) {
	client := &Client{
		UserID: UserID,
		Conn:   conn,
	}
	m.clientsMux.Lock()
	defer m.clientsMux.Unlock()
	m.clients[client.UserID] = client
}

func (m *Manager) RemoveClient(UserID uuid.UUID) {
	m.clientsMux.Lock()
	defer m.clientsMux.Unlock()
	delete(m.clients, UserID)
}

func (m *Manager) GetClient(UserID uuid.UUID) *Client {
	m.clientsMux.RLock()
	defer m.clientsMux.RUnlock()
	return m.clients[UserID]
}
