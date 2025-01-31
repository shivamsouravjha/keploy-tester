package redis

import (
	_ "segwise/utils"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

var RedisClient *redis.Client

func RedisSession() *redis.Client {
	return RedisClient
}

func init() {
	// Connect to the Redis server
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password for local Redis, set it if needed
		DB:       0,                // Default DB
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		zap.L().Error("Error connecting to Redis", zap.Error(err))
	}
	zap.L().Info("Connected to Redis")
}
