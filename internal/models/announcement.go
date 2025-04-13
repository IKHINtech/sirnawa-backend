package models

import "time"

type Announcement struct {
	BaseModel
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedBy uint      `json:"created_by"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}
