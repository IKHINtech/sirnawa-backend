package models

type Rt struct {
	BaseModel
	Name          string      `json:"name" gorm:"not null"`
	RwID          string      `json:"rw_id" gorm:"not null"`
	Rw            Rw          `json:"rw" gorm:"foreignKey:RwID"`
	HousingAreaID string      `gorm:"not null" json:"housing_area_id"`
	HousingArea   HousingArea `gorm:"foreignKey:HousingAreaID" json:"housing_area"`
}

type Rts []Rt
