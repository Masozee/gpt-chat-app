package db

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

type RedisDB struct {
	client *redis.Client
}

func NewRedisDB(addr string) (*RedisDB, error) {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	// Ping the Redis server to verify connection
	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	log.Println("Connected to Redis")
	return &RedisDB{client: client}, nil
}

func (r *RedisDB) Close() {
	if r.client != nil {
		if err := r.client.Close(); err != nil {
			log.Printf("Error closing Redis connection: %v", err)
		}
	}
}

// Add methods for Redis operations, e.g.:
func (r *RedisDB) SetUserPresence(ctx context.Context, userID string, status string) error {
	return r.client.Set(ctx, "presence:"+userID, status, 0).Err()
}

func (r *RedisDB) GetUserPresence(ctx context.Context, userID string) (string, error) {
	return r.client.Get(ctx, "presence:"+userID).Result()
}

// Add more methods as needed for your application
