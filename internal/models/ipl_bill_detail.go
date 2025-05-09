package models

type IplBillDetail struct {
	BaseModel
	IplBillID string  `json:"ipl_bill_id" gorm:"not null"`
	IplBill   IplBill `gorm:"foreignKey:IplBillID" json:"ipl_bill"`
	ItemID    string  `json:"item_id" gorm:"not null"`
	Item      Item    `json:"item" gorm:"foreignKey:ItemID"`
	Note      string  `json:"note" gorm:"type:text"`
	SubAmount int64   `json:"sub_amount"`
}

type IplBillDetails []IplBillDetail
