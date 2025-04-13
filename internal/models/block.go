package models

type Block struct {
	BaseModel
	Name   string `gorm:"uniqueIndex" json:"name"`
	Houses Houses `gorm:"foreignKey:BlockID" json:"houses,omitempty"`
}
