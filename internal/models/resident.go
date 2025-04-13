package models

import "time"

type Resident struct {
	BaseModel
	HouseID        uint      `json:"house_id"`
	Name           string    `json:"name"`
	NIK            string    `json:"nik"`
	BirthDate      time.Time `json:"birth_date"`
	Gender         string    `json:"gender"`
	Job            string    `json:"job"`
	IsHeadOfFamily bool      `json:"is_head_of_family"`
	User           *User     `gorm:"foreignKey:ResidentID" json:"user,omitempty"`
}

type Residents []Resident
