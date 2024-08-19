package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/yourusername/gpt-chat-app/internal/config"
	"github.com/yourusername/gpt-chat-app/internal/server"
)

func TestWebSocketConnection(t *testing.T) {
	cfg := &config.Config{
		Port:       8080,
		MaxClients: 300,
		AIProvider: "chatgpt",
	}

	srv := server.NewServer(cfg)

	s := httptest.NewServer(http.HandlerFunc(srv.HandleWebSocket))
	defer s.Close()

	u := "ws" + strings.TrimPrefix(s.URL, "http")

	ws, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer ws.Close()

	// Add more specific tests for WebSocket communication
}
