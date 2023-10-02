package database

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8" // Importing the Redis client library
)

// Ctx is a shared context for Redis operations.
var Ctx = context.Background()

// CreateClient creates a new Redis client and returns it.
func CreateClient(dbNo int) *redis.Client {
	// Create a new Redis client with the specified configuration
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDR"), // Address of the Redis server obtained from environment variables
		Password: os.Getenv("DB_PASS"), // Password for the Redis server obtained from environment variables
		DB:       dbNo,                 // The Redis database number to use (0 or 1 in your case)
	})

	// Return the created Redis client
	return rdb
}
