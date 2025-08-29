// Package websocket provides WebSocket connection management and message broadcasting.
package websocket

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// BroadcastMessage represents a message to broadcast with sender information.
type BroadcastMessage struct {
	data   []byte
	sender *websocket.Conn
}

// Hub manages WebSocket connections and broadcasting.
type Hub struct {
	clients   map[*websocket.Conn]bool
	content   string
	broadcast chan BroadcastMessage
	upgrader  websocket.Upgrader
	mu        sync.RWMutex
}

// NewHub creates a new WebSocket hub.
func NewHub() *Hub {
	return &Hub{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan BroadcastMessage, 256),
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				// TODO: Implement proper origin checking
				return true
			},
		},
	}
}

// Start begins the message broadcasting goroutine.
func (h *Hub) Start() {
	go h.run()
}

// run handles message broadcasting to all connected clients.
func (h *Hub) run() {
	for message := range h.broadcast {
		h.broadcastToClients(message.data, message.sender)
	}
}

// broadcastToClients sends a message to all connected clients except the sender.
func (h *Hub) broadcastToClients(message []byte, sender *websocket.Conn) {
	h.mu.RLock()
	clients := make([]*websocket.Conn, 0, len(h.clients))
	for conn := range h.clients {
		if conn != sender {
			clients = append(clients, conn)
		}
	}
	h.mu.RUnlock()

	for _, conn := range clients {
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Printf("Error writing message to WebSocket: %v", err)
			h.removeClient(conn)
		}
	}
}

// HandleConnection handles a new WebSocket connection.
func (h *Hub) HandleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}

	h.addClient(conn)
	defer h.removeClient(conn)

	// Send current content to new client
	if content := h.GetContent(); content != "" {
		if err := conn.WriteMessage(websocket.TextMessage, []byte(content)); err != nil {
			log.Printf("Error sending initial content: %v", err)
			return
		}
	}

	// Handle incoming messages
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			h.logConnectionError(err)
			break
		}

		h.updateContent(string(message))
		select {
		case h.broadcast <- BroadcastMessage{data: message, sender: conn}:
		default:
			log.Printf("Broadcast channel full, dropping message")
		}
	}
}

// logConnectionError logs WebSocket connection errors appropriately.
func (h *Hub) logConnectionError(err error) {
	if websocket.IsCloseError(err,
		websocket.CloseGoingAway,
		websocket.CloseNormalClosure,
		websocket.CloseNoStatusReceived,
	) {
		log.Printf("WebSocket connection closed normally")
	} else if websocket.IsUnexpectedCloseError(err,
		websocket.CloseGoingAway,
		websocket.CloseAbnormalClosure,
		websocket.CloseNormalClosure,
	) {
		log.Printf("WebSocket unexpected error: %v", err)
	}
}

// GetContent returns the current content safely.
func (h *Hub) GetContent() string {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.content
}

// updateContent updates the stored content safely.
func (h *Hub) updateContent(content string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.content = content
}

// addClient adds a new client connection safely.
func (h *Hub) addClient(conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.clients[conn] = true
	log.Printf("Client connected. Total clients: %d", len(h.clients))
}

// removeClient removes a client connection safely.
func (h *Hub) removeClient(conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if _, exists := h.clients[conn]; exists {
		delete(h.clients, conn)
		conn.Close()
		log.Printf("Client disconnected. Total clients: %d", len(h.clients))
	}
}
