package services

import (
	"context"

	"github.com/zzhunght/realtime-video-ranking/internal/domain/models"
	"github.com/zzhunght/realtime-video-ranking/internal/domain/repositories"
)

type VideoService struct {
	repo repositories.VideoRepository
}

func NewVideoService(repo repositories.VideoRepository) *VideoService {
	return &VideoService{
		repo: repo,
	}
}

func (s *VideoService) GetVideo(ctx context.Context, videoID string) (*models.Video, error) {
	video, err := s.repo.GetByID(ctx, videoID)
	return video, err
}
