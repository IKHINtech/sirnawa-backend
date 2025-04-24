package models

type IplDetail struct {
	BaseModel
	IplID     string  `json:"ipl_id" gorm:"not null"`
	Ipl       Ipl     `gorm:"foreignKey:IplID" json:"ipl"`
	Note      string  `json:"note" gorm:"type:text"`
	SubAmount float64 `json:"sub_amount"`
}
