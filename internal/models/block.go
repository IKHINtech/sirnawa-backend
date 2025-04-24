package models

type Block struct {
	BaseModel
	Name   string `gorm:"uniqueIndex" json:"name"`
	RtID   string `gorm:"not null" json:"rt_id"`
	Rt     Rt     `gorm:"foreignKey:RtID" json:"rt"`
	Houses Houses `gorm:"foreignKey:BlockID" json:"houses,omitempty"`
}

type Blocks []Block

//DISINI
