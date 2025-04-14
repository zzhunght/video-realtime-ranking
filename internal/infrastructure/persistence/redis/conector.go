package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient(host string, port int, password string, db int) (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: password,
		DB:       db,
	})

	// Kiểm tra kết nối
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("could not connect to redis: %w", err)
	}

	return &RedisClient{Client: client}, nil
}
