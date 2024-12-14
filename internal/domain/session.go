package domain

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID           uuid.UUID `json:"id"`
	UserID       uuid.UUID `json:"user_id"`
	AccessToken  uuid.UUID `json:"access_token"`
	RefreshToken uuid.UUID `json:"refresh_token"`
	SignedAt     time.Time `json:"signed_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
