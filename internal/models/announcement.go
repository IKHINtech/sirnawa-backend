package models

import "github.com/lib/pq"

type Announcement struct {
	BaseModel
	Title       string         `json:"title" gorm:"not null"`
	Content     string         `json:"content" gorm:"not null"`
	CreatedBy   string         `json:"created_by" gorm:"not null"`
	User        User           `json:"user" gorm:"foreignKey:CreatedBy"`
	RtID        string         `gorm:"not null"`
	Rt          Rt             `gorm:"foreignKey:RtID"`
	Attachments pq.StringArray `json:"attachments" gorm:"type:text[]"`
}

type Announcements []Announcement
