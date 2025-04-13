package models

type RondaGroup struct {
	BaseModel
	Name    string             `json:"name"`
	Members []RondaGroupMember `gorm:"foreignKey:GroupID" json:"members,omitempty"`
}
