package models

import (
	"github.com/google/uuid"
)

type RondaAttendance struct {
	BaseModel
	RondaActivityID uuid.UUID     `json:"ronda_activity_id"`
	ResidentID      uuid.UUID     `json:"resident_id"`
	Status          string        `json:"status"` // hadir / tidak_hadir
	Note            string        `json:"note"`
	Resident        Resident      `gorm:"foreignKey:ResidentID"`
	RondaActivity   RondaActivity `gorm:"foreignKey:RondaActivityID"`
}
