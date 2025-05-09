package models

import "time"

type IplRate struct {
	BaseModel
	RtID      string     `json:"rt_id" gorm:"not null"`
	BlockID   *string    `json:"block_id" gorm:"null"`
	HouseType *HouseType `json:"house_type" gorm:"null"`
	Amount    int64      `json:"amount" gorm:"not null"`
	StartDate time.Time  `json:"start_date" gorm:"not null"`

	Rt             Rt             `gorm:"foreignKey:RtID" json:"rt"`
	Block          Block          `gorm:"foreignKey:BlockID" json:"block"`
	IplRateDetails IplRateDetails `json:"items"`
}
type IplRates []IplRate
