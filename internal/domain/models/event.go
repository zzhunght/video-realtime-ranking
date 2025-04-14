package models

import (
	"time"
)

type EventType string

const (
	View    EventType = "VIEW"
	React   EventType = "REACT"
	Comment EventType = "COMMENT"
	Share   EventType = "SHARE"
)

type Event struct {
	ID        int64     `json:"id"`
	VideoID   string    `json:"video_id"`
	UserID    string    `json:"user_id"`
	Type      EventType `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}
