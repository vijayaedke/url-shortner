package redis

import (
	"fmt"

	"url-shortner/internal/config"

	"github.com/go-redis/redis"
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient(config config.Configuration) RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.GetRedisHostname(), config.GetRedisPort()),
		Password: config.GetRedisPassword(),
		DB:       config.GetRedisDatabase(),
	})

	return RedisClient{
		Client: rdb,
	}
}
