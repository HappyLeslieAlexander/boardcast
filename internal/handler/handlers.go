// Package handler provides HTTP handlers for the boardcast application.
package handler

import (
	"net/http"

	"github.com/yosebyte/boardcast/internal/auth"
	"github.com/yosebyte/boardcast/internal/template"
	"github.com/yosebyte/boardcast/internal/websocket"
)

// Handlers contains all HTTP handlers for the application.
type Handlers struct {
	auth  *auth.Manager
	wsHub *websocket.Hub
}

// New creates a new Handlers instance.
func New(authManager *auth.Manager, wsHub *websocket.Hub) *Handlers {
	return &Handlers{
		auth:  authManager,
		wsHub: wsHub,
	}
}

// ServeWhiteboard serves the main whiteboard page.
func (h *Handlers) ServeWhiteboard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(template.WhiteboardHTML))
}

// HandleAuth handles authentication requests.
func (h *Handlers) HandleAuth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	h.auth.Login(w, r)
}

// HandleLogout handles logout requests.
func (h *Handlers) HandleLogout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	h.auth.Logout(w, r)
}

// HandleWebSocket handles WebSocket connections with authentication.
func (h *Handlers) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	if !h.auth.IsAuthenticated(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	h.wsHub.HandleConnection(w, r)
}

// HandleContent returns the current whiteboard content with authentication.
func (h *Handlers) HandleContent(w http.ResponseWriter, r *http.Request) {
	if !h.auth.IsAuthenticated(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(h.wsHub.GetContent()))
}

// HandleSave saves the current whiteboard content to a file.
func (h *Handlers) HandleSave(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if !h.auth.IsAuthenticated(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if err := h.wsHub.SaveSnapshot(); err != nil {
		http.Error(w, "Failed to save snapshot", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Snapshot saved successfully"))
}

// HandleRestore restores whiteboard content from a snapshot file.
func (h *Handlers) HandleRestore(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if !h.auth.IsAuthenticated(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if err := h.wsHub.RestoreSnapshot(); err != nil {
		http.Error(w, "Failed to restore snapshot", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Snapshot restored successfully"))
}
