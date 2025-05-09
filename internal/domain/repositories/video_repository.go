package repositories

import (
	"context"

	"github.com/zzhunght/realtime-video-ranking/internal/domain/models"
)

type VideoRepository interface {
	GetByID(ctx context.Context, id string) (*models.Video, error)
	Create(ctx context.Context, video *models.Video) error
	// ... implement thêm các method khác ở đây
}
