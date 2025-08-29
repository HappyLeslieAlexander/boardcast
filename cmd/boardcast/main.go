// Package main is the entry point for the boardcast application.
package main

import (
	"log"

	"github.com/yosebyte/boardcast/internal"
	"github.com/yosebyte/boardcast/internal/config"
)

var version = "dev"

func main() {
	// Load configuration
	cfg := config.Load(version)

	// Create and configure server
	server, err := internal.NewServer(cfg)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	// Start server
	if err := server.Start(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
