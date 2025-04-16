package models

import "time"

type Resident struct {
	BaseModel
	HouseID        string    `gorm:"not null" json:"house_id"`
	Name           string    `gorm:"not null" json:"name"`
	NIK            string    `gorm:"not null" json:"nik"`
	PhoneNumber    *string   `gorm:"null" json:"phone_number"`
	BirthDate      time.Time `gorm:"not null" json:"birth_date"`
	Gender         string    `gorm:"not null" json:"gender"`
	Job            string    `gorm:"not null" json:"job"`
	IsHeadOfFamily bool      `gorm:"default:false" json:"is_head_of_family"`
	User           *User     `gorm:"foreignKey:ResidentID" json:"user,omitempty"`
	House          House     `gorm:"foreignKey:HouseID" json:"house"`
}

type Residents []Resident
