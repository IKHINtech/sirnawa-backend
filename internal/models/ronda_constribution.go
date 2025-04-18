package models

import "time"

type RondaConstribution struct {
	BaseModel
	Date            time.Time     `json:"date"`
	RondaGroupID    string        `json:"ronda_group_id"`
	RondaActivityID string        `json:"ronda_activity_id" gprm:"not null"`
	RondaGroup      RondaGroup    `gorm:"foreignKey:RondaGroupID" json:"ronda_group"`
	RondaActivity   RondaActivity `json:"ronda_activity" gorm:"foreignKey:RondaActivityID"`
	Total           float64       `json:"total"`
	TotalPenalty    float64       `json:"total_penalty"`
}

type RondaConstributions []RondaConstribution
