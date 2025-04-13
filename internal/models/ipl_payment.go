package models

import "time"

type IplPayment struct {
	BaseModel
	HouseID string           `json:"house_id"`
	House   House            `gorm:"foreignKey:HouseID" json:"house"`
	Month   int              `gorm:"not null" json:"month"`
	Year    int              `gorm:"not null" json:"year"`
	Amount  float64          `gorm:"not null" json:"amount"`
	Status  IplPaymentStatus `gorm:"default:unpaid" json:"status"` // paid/unpaid
	PaidAt  *time.Time       `json:"paid_at,omitempty"`
}
