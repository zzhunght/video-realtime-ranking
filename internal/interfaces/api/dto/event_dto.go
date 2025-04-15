package dto

type AddScore struct {
	UserID    string                 `json:"user_id"`
	VideoID   string                 `json:"video_id"`
	EventType string                 `json:"type"`
	Payload   map[string]interface{} `json:"payload"`
}
