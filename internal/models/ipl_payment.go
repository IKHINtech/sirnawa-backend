package models

import "time"

type IplPayment struct {
	BaseModel
	HouseID string           `gorm:"not null" json:"house_id"`
	House   House            `gorm:"foreignKey:HouseID" json:"house"`
	Month   int              `gorm:"not null" json:"month"`
	Year    int              `gorm:"not null" json:"year"`
	Amount  float64          `gorm:"not null" json:"amount"`
	Status  IplPaymentStatus `gorm:"default:unpaid;type:ipl_payment_status" json:"status"` // paid/unpaid
	PaidAt  *time.Time       `gorm:"null" json:"paid_at,omitempty"`
}

type IplPayments []IplPayment
