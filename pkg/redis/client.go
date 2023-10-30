package redis

import (
	"fmt"
	"time"

	"url-shortner/internal/config"

	"github.com/go-redis/redis"
)

type RedisClient struct {
	Client *redis.Client
}

type RedisService interface {
	SetURLStore(key, value string) (string, error)
	GetURLStore(key string) (string, error)
	GetAllURLStore() ([]string, error)
}

func NewRedisClient(config config.Configuration) RedisService {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.GetRedisHostname(), config.GetRedisPort()),
		Password: config.GetRedisPassword(),
		DB:       config.GetRedisDatabase(),
	})

	return &RedisClient{
		Client: rdb,
	}
}

func (r *RedisClient) SetURLStore(key, value string) (string, error) {
	duration := time.Duration(1 * time.Hour)

	setURL, err := r.Client.SetNX(key, value, duration).Result()
	if err != nil {
		return "", err
	}

	if !setURL {
		fmt.Println("URL already shortened")
		return r.GetURLStore(key)
	}
	return "", nil
}

func (r *RedisClient) GetURLStore(key string) (string, error) {
	result, err := r.Client.Get(key).Result()
	if err != nil {
		return "", err
	}
	return result, err
}

func (r *RedisClient) GetAllURLStore() ([]string, error) {
	var cursor uint64
	var keys, result []string
	var err error
	for {
		keys, cursor, err = r.Client.Scan(cursor, "*", 0).Result()
		if err != nil {
			return nil, err
		}
		result = append(result, keys...)
		if cursor == 0 { // no more keys
			return result, nil
		}
	}
	return nil, fmt.Errorf("no urls found")
}
