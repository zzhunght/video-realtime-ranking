package dto

import (
	"errors"

	"github.com/zzhunght/realtime-video-ranking/internal/domain/models"
)

type AddScore struct {
	UserID    string                 `json:"user_id"`
	VideoID   string                 `json:"video_id"`
	EventType string                 `json:"type"`
	Payload   map[string]interface{} `json:"payload"`
}

func (m *AddScore) Validate() error {

	if m.EventType != string(models.ViewEvent) &&
		m.EventType != string(models.ShareEvent) &&
		m.EventType != string(models.ReactEvent) &&
		m.EventType != string(models.CommentEvent) {
		return errors.New("event type must be one of: VIEW, SHARE, REACT, COMMENT")
	}

	if m.VideoID == "" {
		return errors.New("resource not found")
	}

	if m.UserID == "" {
		return errors.New("missing user_id")
	}

	if m.Payload == nil {
		return errors.New("missing payload")
	}

	return nil
}
