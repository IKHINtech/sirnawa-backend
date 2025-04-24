package models

import "time"

type RondaSchedule struct {
	BaseModel
	RtID    string     `gorm:"not null" json:"rt_id"`
	Rt      Rt         `gorm:"foreignKey:RtID" json:"rt"`
	Date    time.Time  `gorm:"not null" json:"date"`
	GroupID string     `gorm:"not null" json:"group_id"`
	Group   RondaGroup `gorm:"foreignKey:GroupID" json:"group"`
}

type RondaSchedules []RondaSchedule

//DISINI
