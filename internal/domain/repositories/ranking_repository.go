package repositories

import (
	"context"

	"github.com/zzhunght/realtime-video-ranking/internal/domain/models"
)

type RankingRepository interface {
	AddScore(ctx context.Context, event *models.Event) error
	GetRanking(ctx context.Context) error
}
