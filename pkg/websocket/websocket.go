package websocket

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Note: In production, implement proper origin checking
	},
}

func Upgrade(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (*websocket.Conn, error) {
	return upgrader.Upgrade(w, r, responseHeader)
}

// Re-export necessary constants and functions from gorilla/websocket
var (
	CloseMessage         = websocket.CloseMessage
	CloseGoingAway       = websocket.CloseGoingAway
	CloseAbnormalClosure = websocket.CloseAbnormalClosure
	TextMessage          = websocket.TextMessage
)

func IsUnexpectedCloseError(err error, expectedCodes ...int) bool {
	return websocket.IsUnexpectedCloseError(err, expectedCodes...)
}
