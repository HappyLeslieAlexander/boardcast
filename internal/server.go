// Package internal provides the core server functionality for the boardcast application.
package internal

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yosebyte/boardcast/internal/auth"
	"github.com/yosebyte/boardcast/internal/config"
	"github.com/yosebyte/boardcast/internal/handler"
	"github.com/yosebyte/boardcast/internal/websocket"
)

// Server represents the main application server.
type Server struct {
	config   *config.Config
	auth     *auth.Manager
	wsHub    *websocket.Hub
	handlers *handler.Handlers
	server   *http.Server
}

// NewServer creates a new server instance with the given configuration.
func NewServer(cfg *config.Config) (*Server, error) {
	// Initialize authentication manager
	authManager, err := auth.NewManager(cfg.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to create auth manager: %w", err)
	}

	// Initialize WebSocket hub
	wsHub := websocket.NewHub()

	// Initialize handlers
	handlers := handler.New(authManager, wsHub)

	// Create HTTP server
	server := &http.Server{
		Addr:         cfg.GetAddr(),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	s := &Server{
		config:   cfg,
		auth:     authManager,
		wsHub:    wsHub,
		handlers: handlers,
		server:   server,
	}

	// Register routes
	s.registerRoutes()

	return s, nil
}

// Start starts the WebSocket hub and HTTP server with graceful shutdown.
func (s *Server) Start() error {
	// Start WebSocket hub
	s.wsHub.Start()

	// Channel to listen for interrupt signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Start server in a goroutine
	go func() {
		log.Printf("Server started: %s | Password: %s", s.server.Addr, s.config.Password)
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal
	<-stop
	log.Println("Shutting down server...")

	// Create context with timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown server gracefully
	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("server shutdown failed: %w", err)
	}

	log.Println("Server stopped")
	return nil
}

// registerRoutes sets up all HTTP routes.
func (s *Server) registerRoutes() {
	http.HandleFunc("/", s.handlers.ServeWhiteboard)
	http.HandleFunc("/auth", s.handlers.HandleAuth)
	http.HandleFunc("/logout", s.handlers.HandleLogout)
	http.HandleFunc("/ws", s.handlers.HandleWebSocket)
	http.HandleFunc("/content", s.handlers.HandleContent)
}
