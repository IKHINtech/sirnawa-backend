package models

type House struct {
	BaseModel
	BlockID   string      `gorm:"not null" json:"block_id"`
	RtID      string      `gorm:"not null" json:"rt_id"`
	RwID      string      `gorm:"not null" json:"rw_id"`
	Block     Block       `gorm:"foreignKey:BlockID" json:"block"`
	Rt        Rt          `gorm:"foreignKey:RtID" json:"rt"`
	Rw        Rw          `gorm:"foreignKey:RwID" json:"rw"`
	Number    string      `gorm:"not null" json:"number"`
	Status    HouseStatus `grom:"not null;type:house_status"  json:"status"` // aktif / tidak aktif
	Residents Residents   `gorm:"foreignKey:HouseID" json:"residents,omitempty"`
}

type Houses []House
