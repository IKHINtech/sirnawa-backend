package models

type HousingArea struct {
	BaseModel
	Name      string  `json:"name" gorm:"not null"`
	Latitude  float64 `gorm:"type:decimal(10,8);default:null" json:"latitude,omitempty"`
	Longitude float64 `gorm:"type:decimal(11,8);default:null" json:"longitude,omitempty"`
}

type HousingAreas []HousingArea
