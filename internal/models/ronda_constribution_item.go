package models

import "time"

type RondaContributionItem struct {
	BaseModel
	HouseID   string     `json:"house_id"`
	House     House      `gorm:"foreignKey:HouseID" json:"house"`
	IsPenalty bool       `gorm:"default:false" json:"is_penalty"`
	Amount    float64    `json:"amount"`
	PaidAt    *time.Time `json:"paid_at,omitempty"`
	Notes     string     `json:"notes,omitempty"`
}
