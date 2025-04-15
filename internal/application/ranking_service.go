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

	scores, err := s.rankRepo.GetVideoRanking(ctx, limit, reverse)

	if err == nil {

	}
	for _, score := range scores {
		fmt.Println(score.Member, score.Score)
		metadata, err := s.cachedRepo.HGetAll(ctx, score.Member)

		// nếu không có video trong cache thì lấy từ db -> lưu lại vào cached
		if err != nil {
			video, err := s.videoRepo.GetByID(ctx, score.Member)
			if err != nil {
				fmt.Println("Error getting video from DB:", err)
				fmt.Printf("skip this video : %v \n", score.Member)
				continue
			}

			// lưu lại metadata vào cached
			s.cachedRepo.HSet(
				ctx,
				score.Member,
				[]string{"title", video.Title, "desc", video.Desc},
			)
		}

		fmt.Println("Metadata:", metadata)
	}

	fmt.Println("", scores)
	return nil, nil
}
