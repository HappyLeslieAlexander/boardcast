// Package config provides configuration management for the boardcast application.
package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Config holds all configuration values for the application.
type Config struct {
	Port     string
	Password string
	Version  string
}

// Load parses command line flags and returns a validated Config instance.
func Load(version string) *Config {
	var (
		port        = flag.String("port", "8080", "Port to run the server on (1-65535)")
		password    = flag.String("password", "", "Password for authentication (required)")
		versionFlag = flag.Bool("version", false, "Show version and exit")
	)
	flag.Parse()

	// Handle version flag
	if *versionFlag {
		fmt.Printf("boardcast version %s\n", version)
		os.Exit(0)
	}

	cfg := &Config{
		Port:     *port,
		Password: *password,
		Version:  version,
	}

	if err := cfg.validate(); err != nil {
		log.Fatal(err)
	}

	return cfg
}

// validate checks if the configuration values are valid.
func (c *Config) validate() error {
	if c.Password == "" {
		return fmt.Errorf("password is required. Use --password flag to set it")
	}

	if port, err := strconv.Atoi(c.Port); err != nil || port < 1 || port > 65535 {
		return fmt.Errorf("invalid port number: %s (must be 1-65535)", c.Port)
	}

	return nil
}

// GetAddr returns the formatted address string for the server.
func (c *Config) GetAddr() string {
	return ":" + c.Port
}
