package models

import "time"

type RondaConstribution struct {
	BaseModel
	Date         time.Time  `json:"date"`
	RondaGroupID string     `json:"ronda_group_id"`
	RondaGroup   RondaGroup `gorm:"foreignKey:RondaGroupID" json:"ronda_group"`
	Total        float64    `json:"total"`
	TotalPenalty float64    `json:"total_penalty"`
}
