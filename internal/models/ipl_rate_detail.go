package models

type IplRateDetail struct {
	BaseModel
	IplRateID string `json:"ipl_rate_id" gorm:"not null"`
	ItemID    string `json:"item_id" gorm:"not null"`
	Amount    int64  `json:"amount" gorm:"not null"`

	IplRate IplRate `gorm:"foreignKey:IplRateID" json:"ipl_rate"`
	Item    Item    `gorm:"foreignKey:ItemID" json:"item"`
}

type IplRateDetails []IplRateDetail
