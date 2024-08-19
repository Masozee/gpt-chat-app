package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       int
	MaxClients int
	AIProvider string
	MongoDBURI string
	RedisAddr  string
}

func Load() (*Config, error) {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	maxClients, _ := strconv.Atoi(os.Getenv("MAX_CLIENTS"))

	return &Config{
		Port:       port,
		MaxClients: maxClients,
		AIProvider: os.Getenv("AI_PROVIDER"),
		MongoDBURI: os.Getenv("MONGODB_URI"),
		RedisAddr:  os.Getenv("REDIS_ADDR"),
	}, nil
}
