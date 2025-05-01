package models

import (
	"time"
)

type UserVerification struct {
	BaseModel
	UserID    string    `gorm:"type:uuid;not null" json:"user_id"`
	Code      string    `gorm:"type:varchar(10);not null" json:"code"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`
	IsUsed    bool      `gorm:"default:false" json:"is_used"`
}
