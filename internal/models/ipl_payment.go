package models

import "time"

type IPLPayment struct {
	BaseModel
	HouseID uint       `json:"house_id"`
	House   House      `gorm:"foreignKey:HouseID" json:"house"`
	Month   int        `json:"month"`
	Year    int        `json:"year"`
	Amount  float64    `json:"amount"`
	Status  string     `json:"status"` // paid/unpaid
	PaidAt  *time.Time `json:"paid_at,omitempty"`
}
