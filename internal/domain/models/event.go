package models

import (
	"time"
)

type EventType string

const (
	ViewEvent    EventType = "VIEW"
	ReactEvent   EventType = "REACT"
	CommentEvent EventType = "COMMENT"
	ShareEvent   EventType = "SHARE"
)

type Event struct {
	ID        int64                  `json:"id"`
	VideoID   string                 `json:"video_id"`
	UserID    string                 `json:"user_id"`
	Type      EventType              `json:"type"`
	Payload   map[string]interface{} `json:"payload"`
	CreatedAt time.Time              `json:"created_at"`
}
