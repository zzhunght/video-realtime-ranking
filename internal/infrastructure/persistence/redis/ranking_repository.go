package redis

import "github.com/go-redis/redis/v8"

type RedisRankingRepository struct {
	client *redis.Client
}

func NewRedisRankingRepository(client *redis.Client) *RedisRankingRepository {
	return &RedisRankingRepository{
		client: client,
	}
}

func AddScore() error {

	return nil
}
