package database

import (
	"context"
	CONFIG "video-api/config"

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

func GetFromRedisSortedSets(key, min, max string, offset, count int64) ([]string, error) {
	return redisClient.ZRangeByScore(context.TODO(), key, &redis.ZRangeBy{
		Min:    min,
		Max:    max,
		Offset: offset,
		Count:  count,
	}).Result()
}

func AddToRedisSortedSets(key string, score float64, member string) (int64, error) {
	return redisClient.ZAdd(context.TODO(), key, &redis.Z{
		Score:  score,
		Member: member,
	}).Result()
}

func GetRedisValue(key string) (string, error) {
	return redisClient.Get(context.TODO(), key).Result()
}

func SetRedisValue(key, value string) {
	redisClient.Set(context.TODO(), key, value, 0)
}
