package services

import (
	"context"
	"fmt"

	"github.com/zzhunght/realtime-video-ranking/internal/domain/models"
	"github.com/zzhunght/realtime-video-ranking/internal/domain/repositories"
)

type RankingService struct {
	rankRepo   repositories.RankingRepository
	cachedRepo repositories.CachedRepository
	videoRepo  repositories.VideoRepository
}

func NewRankingService(
	rankRepo repositories.RankingRepository,
	cachedRepo repositories.CachedRepository,
	videoRepo repositories.VideoRepository,
) *RankingService {
	return &RankingService{
		rankRepo:   rankRepo,
		cachedRepo: cachedRepo,
		videoRepo:  videoRepo,
	}
}

func (s *RankingService) GetVideoByRank(ctx context.Context, limit int, reverse bool) ([]models.VideoRank, error) {

	var data []models.VideoRank
	scores, err := s.rankRepo.GetVideoRanking(ctx, limit, reverse)

	if err != nil {
		return nil, err
	}

	for _, score := range scores {
		metadata, err := s.cachedRepo.HGetAll(ctx, score.Member)

		// nếu không có video trong cache thì lấy từ db -> lưu lại vào cached
		if err != nil || len(metadata) == 0 {
			video, err := s.videoRepo.GetByID(ctx, score.Member)
			if err != nil {
				fmt.Println("Error getting video from DB:", err)
				fmt.Printf("skip this video : %v \n", score.Member)
				continue
			}
			_ = s.cachedRepo.HSet(ctx, score.Member, map[string]string{
				"title": video.Title,
				"desc":  video.Desc,
			})

			data = append(data, models.VideoRank{
				ID:    score.Member,
				Score: score.Score,
				Title: video.Title,
			})
			continue
		}

		// Có sẵn trong cache
		data = append(data, models.VideoRank{
			ID:    score.Member,
			Score: score.Score,
			Title: metadata["title"],
		})

	}

	return data, nil
}

func (s *RankingService) AddScore(ctx context.Context, member string, score float64) error {
	return s.rankRepo.AddVideoScore(ctx, member, score)
}
