package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/yourusername/gpt-chat-app/internal/config"
	"github.com/yourusername/gpt-chat-app/internal/db"
)

type Server struct {
	config     *config.Config
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	mutex      sync.Mutex
	mongoDB    *db.MongoDB
	redisDB    *db.RedisDB
}

func NewServer(cfg *config.Config) (*Server, error) {
	mongoDB, err := db.NewMongoDB(cfg.MongoDBURI)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	redisDB, err := db.NewRedisDB(cfg.RedisAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %v", err)
	}

	return &Server{
		config:     cfg,
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		mongoDB:    mongoDB,
		redisDB:    redisDB,
	}, nil
}

func (s *Server) Run() error {
	defer s.mongoDB.Close()
	defer s.redisDB.Close()

	http.HandleFunc("/ws", s.handleWebSocket)

	go s.handleMessages()

	addr := fmt.Sprintf(":%d", s.config.Port)
	log.Printf("Server starting on %s", addr)
	return http.ListenAndServe(addr, nil)
}

// ... (rest of the Server methods remain the same)

func (s *Server) handleMessages() {
	for {
		select {
		case client := <-s.register:
			s.mutex.Lock()
			s.clients[client] = true
			s.mutex.Unlock()
			// Set user presence in Redis
			err := s.redisDB.SetUserPresence(client.ctx, client.ID, "online")
			if err != nil {
				log.Printf("Error setting user presence: %v", err)
			}
		case client := <-s.unregister:
			s.mutex.Lock()
			if _, ok := s.clients[client]; ok {
				delete(s.clients, client)
				close(client.send)
			}
			s.mutex.Unlock()
			// Set user presence in Redis
			err := s.redisDB.SetUserPresence(client.ctx, client.ID, "offline")
			if err != nil {
				log.Printf("Error setting user presence: %v", err)
			}
		case message := <-s.broadcast:
			// Save message to MongoDB
			err := s.mongoDB.SaveMessage(context.Background(), message)
			if err != nil {
				log.Printf("Error saving message to MongoDB: %v", err)
			}
			for client := range s.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(s.clients, client)
				}
			}
		}
	}
}
