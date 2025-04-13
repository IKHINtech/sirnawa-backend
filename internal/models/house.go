package models

type House struct {
	BaseModel
	BlockID   string      `gorm:"not null" json:"block_id"`
	Block     Block       `gorm:"foreignKey:BlockID" json:"block"`
	Number    string      `gorm:"not null" json:"number"`
	Status    HouseStatus `grom:"not null"  json:"status"` // aktif / tidak aktif
	Residents Residents   `gorm:"foreignKey:HouseID" json:"residents,omitempty"`
}

type Houses []House
