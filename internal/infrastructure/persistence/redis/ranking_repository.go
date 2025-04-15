package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/zzhunght/realtime-video-ranking/internal/domain/models"
	"github.com/zzhunght/realtime-video-ranking/internal/domain/repositories"
)

type RedisRankingRepository struct {
	client *redis.Client
}

func NewRedisRankingRepository(client *redis.Client) *RedisRankingRepository {
	return &RedisRankingRepository{
		client: client,
	}
}

func (r *RedisRankingRepository) AddVideoScore(ctx context.Context, member string, score float64) error {
	return r.client.ZIncrBy(ctx, VideoLeaderBoardKey, score, member).Err()
}

func (r *RedisRankingRepository) GetVideoRanking(ctx context.Context, limit int, reverse bool) ([]models.Score, error) {
	var videos []redis.Z
	var err error
	if reverse {
		videos, err = r.client.ZRevRangeWithScores(ctx, VideoLeaderBoardKey, 0, int64(limit)).Result()

	} else {
		videos, err = r.client.ZRangeWithScores(ctx, VideoLeaderBoardKey, 0, int64(limit)).Result()
	}

	if err != nil {
		return nil, err
	}

	var result []models.Score
	for _, video := range videos {
		result = append(result, models.Score{
			Member: video.Member.(string),
			Score:  video.Score,
		})
	}

	return result, nil
}

var _ repositories.RankingRepository = (*RedisRankingRepository)(nil)
