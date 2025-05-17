package models

import (
	"time"

	"gorm.io/gorm"
)

type UserFCMToken struct {
	ID         uint64    `gorm:"primaryKey"`
	UserID     string    `gorm:"not null"`
	Token      string    `gorm:"type:text;not null"`
	DeviceID   string    `gorm:"size:255"`
	DeviceType string    `gorm:"size:50;check:device_type IN ('android','ios','web','other')"`
	AppVersion string    `gorm:"size:50"`
	OSVersion  string    `gorm:"size:50"`
	IsActive   bool      `gorm:"default:true"`
	CreatedAt  time.Time `gorm:"not null;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"not null;autoUpdateTime"`
	ExpiresAt  *time.Time

	User User `gorm:"foreignKey:UserID"`
}

// BeforeCreate hook untuk set expires_at
func (t *UserFCMToken) BeforeCreate(tx *gorm.DB) error {
	expiry := time.Now().Add(6 * 30 * 24 * time.Hour) // 6 bulan
	t.ExpiresAt = &expiry
	return nil
}
