package test

import (
	"testing"

	"github.com/yourusername/gpt-chat-app/internal/config"
	"github.com/yourusername/gpt-chat-app/internal/server"
)

func TestNewServer(t *testing.T) {
	cfg := &config.Config{
		Port:       8080,
		MaxClients: 300,
		AIProvider: "chatgpt",
	}

	srv := server.NewServer(cfg)

	if srv == nil {
		t.Error("Expected non-nil server")
	}

	// Add more specific tests as needed
}
