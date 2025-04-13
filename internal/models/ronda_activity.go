package models

import "time"

type RondaActivity struct {
	BaseModel
	RondaGroupID string     `json:"ronda_group_id"`
	RondaGroup   RondaGroup `gorm:"foreignKey:RondaGroupID" json:"ronda_group"`
	Date         time.Time  `json:"date"`
	Description  string     `json:"description"`
	CreatedBy    string     `json:"created_by"`
	User         User       `json:"user" gorm:"foreignKey:CreatedBy"`
}
