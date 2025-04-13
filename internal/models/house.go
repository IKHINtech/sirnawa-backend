package models

type House struct {
	BaseModel
	BlockID   uint       `json:"block_id"`
	Block     Block      `gorm:"foreignKey:BlockID" json:"block"`
	Number    string     `json:"number"`
	Status    string     `json:"status"` // aktif / tidak aktif
	Residents []Resident `gorm:"foreignKey:HouseID" json:"residents,omitempty"`
}
