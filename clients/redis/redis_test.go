package redis

import (
	"testing"

	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

func TestRedisConnection(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Use your Redis address
		Password: "",               // No password for local Redis
		DB:       0,                // Default database
	})

	// Check if Redis is reachable
	_, err := client.Ping().Result()
	assert.NoError(t, err, "Redis should be reachable")

	// Ensure the client is not nil
	assert.NotNil(t, client, "Redis client should not be nil")
}

// TestRedisSession verifies that RedisSession() returns a valid client
func TestRedisSession(t *testing.T) {
	client := RedisSession()
	assert.NotNil(t, client, "RedisSession should return a valid Redis client")
}
