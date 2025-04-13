package models

import "time"

type RondaContribution struct {
	BaseModel
	HouseID   uint       `json:"house_id"`
	House     House      `gorm:"foreignKey:HouseID" json:"house"`
	IsPenalty bool       `gorm:"default:false" json:"is_penalty"`
	Week      int        `json:"week"`
	Year      int        `json:"year"`
	Amount    float64    `json:"amount"`
	PaidAt    *time.Time `json:"paid_at,omitempty"`
	Notes     string     `json:"notes,omitempty"`
}
