package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/zzhunght/realtime-video-ranking/internal/domain/repositories"
)

type RedisCachedRepository struct {
	client *redis.Client
}

func NewRedisCachedRepository(client *redis.Client) *RedisCachedRepository {
	return &RedisCachedRepository{
		client: client,
	}
}

func (r *RedisCachedRepository) HSet(ctx context.Context, key string, value interface{}) error {
	err := r.client.HSet(ctx, key, value).Err()
	return err
}

func (r *RedisCachedRepository) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	val, err := r.client.HGetAll(ctx, key).Result()
	if err != nil {
		return map[string]string{}, err
	}
	return val, nil
}

var _ repositories.CachedRepository = (*RedisCachedRepository)(nil)
