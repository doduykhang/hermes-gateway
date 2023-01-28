package config

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func getRedisConn(redis Redis) string {
	return fmt.Sprintf("%s:%s", redis.Host, redis.Port)
}

func NewRedis(config *Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
        	Addr:     getRedisConn(config.Redis),
        	Password: config.Redis.Password, // no password set
        	DB:       0,  // use default DB
    	})
	return rdb
}
