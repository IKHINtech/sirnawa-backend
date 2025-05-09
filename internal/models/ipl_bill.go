package models

import "time"

type IplBill struct {
	BaseModel
	HouseID     string        `json:"house_id" gorm:"not null"`
	Month       int           `json:"month" gorm:"not null"`
	Year        int           `json:"year" gorm:"not null"`
	TotalAmount int64         `json:"total_amount" gorm:"not null"`
	AmountPaid  *int64        `json:"amount_paid" gorm:"null"`
	BalanceDue  *int64        `json:"balance_due" gorm:"null"`
	IplRateID   string        `json:"ipl_rate_id" gorm:"not null"`
	IplRate     IplRate       `gorm:"foreignKey:IplRateID" json:"ipl_rate"`
	Status      IplBillStatus `json:"status" gorm:"default:unpaid"`
	DueDate     time.Time     `json:"due_date" gorm:"not null"`
	Penalty     *string       `json:"penalty" gorm:"null"`

	House House `gorm:"foreignKey:HouseID" json:"house"`
}
