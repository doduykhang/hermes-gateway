package config

import (
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func getRedisConn(redis Redis) string {
	return fmt.Sprintf("%s:%s", redis.Host, redis.Port)
}

func NewRedis(config *Config) *redis.Client {
	log.Println("redis conn", getRedisConn(config.Redis))
	rdb := redis.NewClient(&redis.Options{
        	Addr:     getRedisConn(config.Redis),
        	Password: config.Redis.Password, // no password set
        	DB:       0,  // use default DB
    	})
	return rdb
}
