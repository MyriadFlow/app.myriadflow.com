package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Avatar struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	AvatarID   int       `json:"avatar_id"`
	URL        string    `json:"url"`
	UserID     string    `json:"user_id"`
	PhygitalID string    `json:"phygital_id"`
	CreatedAt  time.Time `gorm:"type:timestamp;default:current_timestamp" json:"created_at"`
	UpdatedAt  time.Time `gorm:"type:timestamp;default:current_timestamp" json:"updated_at"`
}

func (a *Avatar) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New()
	return
}