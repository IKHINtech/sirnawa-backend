package models

type IplBillDetail struct {
	BaseModel
	IplBillID string `json:"ipl_bill_id" gorm:"not null"`
	ItemID    string `json:"item_id" gorm:"not null"`
	Note      string `json:"note" gorm:"type:text"`
	SubAmount int64  `json:"sub_amount"`

	IplBill IplBill `gorm:"foreignKey:IplBillID" json:"ipl_bill"`
	Item    Item    `json:"item" gorm:"foreignKey:ItemID"`
}

type IplBillDetails []IplBillDetail
