package models

import "time"

// Enum types
type RoleName string

const (
	RoleViewer  RoleName = "VIEWER"
	RoleCreator RoleName = "CREATOR"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Avatar    string    `json:"avatar,omitempty"`
	Role      RoleName  `json:"role"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
