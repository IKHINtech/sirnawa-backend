package models

type Rw struct {
	BaseModel
	Name string `json:"name" gorm:"not null"`
}

type Rws []Rw
