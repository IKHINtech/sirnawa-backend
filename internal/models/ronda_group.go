package models

type RondaGroup struct {
	BaseModel
	Name    string             `json:"name"`
	Order   uint               `gorm:"not null"` // TODO: autoincrement
	Members []RondaGroupMember `gorm:"foreignKey:GroupID" json:"members,omitempty"`
}

type RondaGroups []RondaGroup
