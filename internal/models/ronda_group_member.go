package models

type RondaGroupMember struct {
	BaseModel
	GroupID    string `json:"group_id"`
	ResidentID string `json:"resident_id"`
	HouseID    string `json:"house_id" gorm:"not null"`

	Group    RondaGroup `gorm:"foreignKey:GroupID" json:"-"`
	Resident Resident   `gorm:"foreignKey:ResidentID" json:"-"`
	House    House      `gorm:"foreignKey:HouseID"`
}

type RondaGroupMembers []RondaGroupMember
