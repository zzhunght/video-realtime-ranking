package repositories

import "context"

type CachedRepository interface {
	HSet(ctx context.Context, key string, value interface{}) error
	HGetAll(ctx context.Context, key string) (map[string]string, error)
}
