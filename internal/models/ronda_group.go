package models

type RondaGroup struct {
	BaseModel
	Name    string            `json:"name"`
	RtID    string            `json:"rt_id" gorm:"not null"`
	Rt      Rt                `gorm:"foreignKey:RtID" json:"rt"`
	Order   uint              `gorm:"not null;autoincrement"`
	Members RondaGroupMembers `gorm:"foreignKey:GroupID" json:"members,omitempty"`
}

type RondaGroups []RondaGroup

// DISINI
