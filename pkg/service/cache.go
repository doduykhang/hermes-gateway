package service

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Cache interface {
	Set(key string, value string) error	
	Get(key string) (string, error)
}

type cache struct {
	redisClient *redis.Client
}

func NewCache(redis *redis.Client) Cache {
	return &cache{
		redisClient: redis,
	}
}

func (c *cache) Set(key string, value string) (error) {
	return c.redisClient.Set(context.Background(), key, value, 0).Err()
}
func (c *cache) Get(key string) (string, error) {
	return c.redisClient.Get(context.Background(), key).Result()
}



