package models

type Rw struct {
	BaseModel
	Name          string      `json:"name" gorm:"not null"`
	HousingAreaID string      `gorm:"not null" json:"housing_area_id"`
	HousingArea   HousingArea `gorm:"foreignKey:HousingAreaID" json:"housing_area"`
}

type Rws []Rw
