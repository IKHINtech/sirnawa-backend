package models

type Ipl struct {
	BaseModel
	RtID     string      `json:"rt_id" gorm:"not null"`
	Rt       Rt          `gorm:"foreignKey:RtID" json:"rt"`
	Amount   float64     `json:"amount"`
	IsActive bool        `json:"is_active" gorm:"default:true"`
	Note     string      `json:"note" gorm:"type:text"`
	Items    []IplDetail `gorm:"foreignKey:IplID" json:"items,omitempty"`
}
