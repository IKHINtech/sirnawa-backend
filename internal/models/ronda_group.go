package models

import "time"

type RondaGroup struct {
	BaseModel
	Name         string             `json:"name"`
	ScheduleDate time.Time          `json:"schedule_date"`
	Members      []RondaGroupMember `gorm:"foreignKey:GroupID" json:"members,omitempty"`
}
