package database

import (
	"context"
	CONFIG "video-parser/config"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

// Connect - connect to elastisearch database with given configuration
func ConnectRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     CONFIG.RedisAddress,
		Password: CONFIG.RedisPassword,
		DB:       CONFIG.RedisDB,
	})
}

func GetPopFirstInRedisLists(key string) (string, error) {
	return redisClient.RPop(context.TODO(), key).Result()
}

func AddToRedisLists(key, member string) {
	redisClient.LPush(context.TODO(), key, member)
}

func GetRedisValue(key string) (string, error) {
	return redisClient.Get(context.TODO(), key).Result()
}

func SetRedisValue(key, value string) {
	redisClient.Set(context.TODO(), key, value, 0)
}
