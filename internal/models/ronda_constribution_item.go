package models

import "time"

type RondaContributionItem struct {
	BaseModel
	RondaContributionID string             `gorm:"not null"`
	RondaContribution   RondaConstribution `gorm:"foreignKey:RondaContributionID" json:"ronda_contribution"`
	HouseID             string             `json:"house_id"`
	House               House              `gorm:"foreignKey:HouseID" json:"house"`
	IsPenalty           bool               `gorm:"default:false" json:"is_penalty"`
	Amount              float64            `json:"amount"`
	PaidAt              *time.Time         `json:"paid_at,omitempty"`
	Notes               string             `json:"notes,omitempty"`
}

type RondaContributionItems []RondaContributionItem
