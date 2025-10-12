package database

import "github.com/redis/go-redis/v9"

type RedisCache struct {
	redisClient *redis.Client
}

func NewRedisCache() *RedisCache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &RedisCache{
		redisClient: rdb,
	}
}
