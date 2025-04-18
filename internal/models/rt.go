package models

type Rt struct {
	BaseModel
	Name string `json:"name" gorm:"not null"`
	RwID string `json:"rw_id" gorm:"not null"`
	Rw   Rw     `json:"rw" gorm:"foreignKey:RwID"`
}

type Rts []Rt
