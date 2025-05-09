package models

import "time"

type IplPayment struct {
	BaseModel
	IplBillID          string             `json:"ipl_bill_id" gorm:"not null"`
	PaidAt             time.Time          `gorm:"not null" json:"paid_at,omitempty"`
	AmountPaid         int64              `json:"amount_paid" gorm:"not null"`
	IplPaymentMethod   PaymentMethod      `json:"payment_method" gorm:"not null"`
	Evidence           string             `json:"evidence" gorm:"type:text"`
	VerifiedBy         *string            `json:"verified_by" gorm:"null"`
	VerificationStatus VerificationStatus `json:"verification_status" gorm:"not null;default:pending"`
	Notes              string             `json:"notes" gorm:"type:text"`

	IplBill IplBill `json:"ipl_bill" gorm:"foreignKey:IplBillID"`
}

type IplPayments []IplPayment
