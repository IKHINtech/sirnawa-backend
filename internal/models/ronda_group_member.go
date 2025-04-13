package models

type RondaGroupMember struct {
	BaseModel
	GroupID    uint       `json:"group_id"`
	ResidentID uint       `json:"resident_id"`
	Group      RondaGroup `gorm:"foreignKey:GroupID" json:"-"`
	Resident   Resident   `gorm:"foreignKey:ResidentID" json:"-"`
}
