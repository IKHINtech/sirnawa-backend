package models

import "github.com/lib/pq"

type Post struct {
	BaseModel
	UserID      string         `gorm:"not null" json:"user_id"`
	User        User           `gorm:"foreignKey:UserID" json:"user"`
	RtID        string         `gorm:"not null" json:"rt_id"`
	Rt          Rt             `gorm:"foreignKey:RtID" json:"rt"`
	Title       string         `gorm:"not null" json:"title"`
	Content     string         `gorm:"not null" json:"content"`
	Attachments pq.StringArray `json:"attachments" gorm:"type:text[]"`
}

type Posts []Post
