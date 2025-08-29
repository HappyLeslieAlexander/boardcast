// Package auth provides authentication functionality for the boardcast application.
package auth

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

const (
	SessionName = "boardcast-session"
	AuthKey     = "authenticated"
	SecretKey   = "boardcast-secret-key"
)

// Manager handles authentication operations.
type Manager struct {
	hashedPassword []byte
	store          *sessions.CookieStore
}

// NewManager creates a new authentication manager.
func NewManager(password string) (*Manager, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &Manager{
		hashedPassword: hashedPassword,
		store:          sessions.NewCookieStore([]byte(SecretKey)),
	}, nil
}

// LoginRequest represents the JSON structure for login requests.
type LoginRequest struct {
	Password string `json:"password"`
}

// IsAuthenticated checks if the request is authenticated.
func (m *Manager) IsAuthenticated(r *http.Request) bool {
	session, err := m.store.Get(r, SessionName)
	if err != nil {
		return false
	}

	auth, ok := session.Values[AuthKey].(bool)
	return ok && auth
}

// Login processes authentication requests.
func (m *Manager) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	if err := bcrypt.CompareHashAndPassword(m.hashedPassword, []byte(req.Password)); err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	if err := m.setAuthStatus(w, r, true); err != nil {
		http.Error(w, "Session error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("authenticated"))
}

// Logout processes logout requests.
func (m *Manager) Logout(w http.ResponseWriter, r *http.Request) {
	if err := m.setAuthStatus(w, r, false); err != nil {
		http.Error(w, "Session error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("logged out"))
}

// setAuthStatus sets the authentication status in the session.
func (m *Manager) setAuthStatus(w http.ResponseWriter, r *http.Request, authenticated bool) error {
	session, err := m.store.Get(r, SessionName)
	if err != nil {
		return err
	}

	session.Values[AuthKey] = authenticated
	return session.Save(r, w)
}
