package models

type House struct {
	BaseModel
	BlockID        string      `gorm:"not null" json:"block_id"`
	RtID           string      `gorm:"not null" json:"rt_id"`
	Rt             Rt          `gorm:"foreignKey:RtID" json:"rt"`
	RwID           string      `gorm:"not null" json:"rw_id"`
	HousingAreaID  string      `gorm:"not null" json:"housing_area_id"`
	Block          Block       `gorm:"foreignKey:BlockID" json:"block"`
	Rw             Rw          `gorm:"foreignKey:RwID" json:"rw"`
	HousingArea    HousingArea `gorm:"foreignKey:HousingAreaID" json:"housing_area"`
	Number         string      `gorm:"not null" json:"number"`
	Status         HouseStatus `gorm:"not null;type:house_status"  json:"status"` // aktif / tidak aktif / kontrak
	Latitude       float64     `gorm:"type:decimal(10,8);default:null" json:"latitude,omitempty"`
	Longitude      float64     `gorm:"type:decimal(11,8);default:null" json:"longitude,omitempty"`
	ResidentHouses ResidentHouses
}

type Houses []House

//DISINI
