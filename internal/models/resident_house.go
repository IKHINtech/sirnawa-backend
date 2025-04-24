package models

type ResidentHouse struct {
	BaseModel
	ResidentID string   `gorm:"not null"`
	Resident   Resident `gorm:"foreignKey:ResidentID"`
	HouseID    string   `gorm:"not null"`
	House      House    `gorm:"foreignKey:HouseID"`
	IsPrimary  bool     `gorm:"default:false"` // Apakah ini rumah utama
}

type ResidentHouses []ResidentHouse
