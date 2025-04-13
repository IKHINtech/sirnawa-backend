package models

type RondaGroupMember struct {
	BaseModel
	GroupID    string     `json:"group_id"`
	ResidentID string     `json:"resident_id"`
	Group      RondaGroup `gorm:"foreignKey:GroupID" json:"-"`
	Resident   Resident   `gorm:"foreignKey:ResidentID" json:"-"`
}
