package server

import (
	"log"
	"net/http"

	"github.com/yourusername/gpt-chat-app/pkg/websocket"
)

func (s *Server) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{
		ID:     r.RemoteAddr,
		conn:   conn,
		send:   make(chan []byte, 256),
		server: s,
	}

	s.register <- client

	go client.writePump()
	go client.readPump()
}
