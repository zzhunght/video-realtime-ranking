package repositories

import (
	"context"

	"github.com/zzhunght/realtime-video-ranking/internal/domain/models"
)

type RankingRepository interface {
	AddVideoScore(ctx context.Context, member string, score float64) error
	GetVideoRanking(ctx context.Context, limit int, reverse bool) ([]models.Score, error)
}
