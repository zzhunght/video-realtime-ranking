package models

import "time"

type Video struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	Desc       string    `json:"desc"`
	CategoryID int       `json:"category_id"`
	CreatorID  string    `json:"creator_id"`
	Score      int       `json:"score"`
	UpdatedAt  time.Time `json:"updated_at"`
	CreatedAt  time.Time `json:"created_at"`
}
