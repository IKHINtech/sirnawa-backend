package models

type HousingArea struct {
	BaseModel
	Name string `json:"name" gorm:"not null"`
}

type HousingAreas []HousingArea
