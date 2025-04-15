package postgres

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/zzhunght/realtime-video-ranking/internal/domain/models"
	"github.com/zzhunght/realtime-video-ranking/internal/domain/repositories"
	errors "github.com/zzhunght/realtime-video-ranking/pkg"
)

type PostgresVideoRepository struct {
	db *sql.DB
}

func NewVideoRepository(db *sql.DB) *PostgresVideoRepository {
	return &PostgresVideoRepository{db: db}
}

func (v *PostgresVideoRepository) GetByID(ctx context.Context, id string) (*models.Video, error) {
	var video models.Video

	row, err := v.db.QueryContext(ctx,
		"SELECT id, title, desc, category_id from video where id = $1;",
		id,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrNotfound
		}

		return nil, err
	}
	row.Scan(
		&video.ID,
		&video.Title,
		&video.Desc,
		&video.CategoryID,
	)

	return &video, nil
}

func (v *PostgresVideoRepository) Create(ctx context.Context, video *models.Video) error {

	query := `
	INSERT INTO videos (title, desc, category_id, creator_id)
	VALUES ($1, $2, $3, $4)
	RETURNING id
	`

	row, err := v.db.QueryContext(
		ctx,
		query,
		video.Title,
		video.Desc,
		video.CategoryID,
		video.CreatorID,
	)

	if err != nil {
		return err
	}
	row.Scan(&video.ID)
	return nil
}

// kiểm tra impl có đủ method của interface hay k
var _ repositories.VideoRepository = (*PostgresVideoRepository)(nil)
