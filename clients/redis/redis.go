package redis

import (
	"segwise/config"
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
		Addr:     config.Get().RedisAddr,     // Redis server address
		Password: config.Get().RedisPassword, // No password for local Redis, set it if needed
		DB:       0,                          // Default DB
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		zap.L().Error("Error connecting to Redis", zap.Error(err))
		return
	}
	zap.L().Info("Connected to Redis")
}
