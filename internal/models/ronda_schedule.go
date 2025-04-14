package models

import "time"

type RondaSchedule struct {
	BaseModel
	EfektifDate time.Time  `gorm:"not null" json:"efektif_date"`
	GroupID     string     `gorm:"not null" json:"group_id"`
	Group       RondaGroup `gorm:"foreignKey:GroupID" json:"group"`
}
