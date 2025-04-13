package models

import "github.com/lib/pq"

type Post struct {
	BaseModel
	UserID      string         `json:"user_id"`
	User        User           `gorm:"foreignKey:UserID" json:"user"`
	Title       string         `gorm:"not null" json:"title"`
	Content     string         `gorm:"not null" json:"content"`
	Attachments pq.StringArray `json:"attachments" gorm:"type:text[]"`
}
