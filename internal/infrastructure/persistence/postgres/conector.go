package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/zzhunght/realtime-video-ranking/internal/config"
)

func NewDB(cfg *config.Config) (*sql.DB, error) {

	db, err := sql.Open("postgres", cfg.DB.DNS)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.DB.MaxOpenConns)
	db.SetMaxIdleConns(cfg.DB.MaxIdleConns)
	duration, err := time.ParseDuration(cfg.DB.MaxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
