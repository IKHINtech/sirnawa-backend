package models

import "time"

type RondaActivity struct {
	BaseModel
	RondaGroupID string     `json:"ronda_group_id"`
	RondaGroup   RondaGroup `gorm:"foreignKey:RondaGroupID" json:"ronda_group"`
	Date         time.Time  `gorm:"not null" json:"date"`
	Description  string     `gorm:"type:text" json:"description"`
	CreatedBy    string     `gorm:"not null" json:"created_by"`
	User         User       `gorm:"foreignKey:CreatedBy" json:"user"`
}
