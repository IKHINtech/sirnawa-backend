package models

type RondaAttendance struct {
	BaseModel
	RondaActivityID string                `json:"ronda_activity_id"`
	ResidentID      string                `json:"resident_id"`
	Status          RondaAttendanceStatus `gorm:"not null;type:ronda_attendance_status" json:"status"` // hadir / tidak_hadir
	Note            string                `json:"note"`
	Resident        Resident              `gorm:"foreignKey:ResidentID"`
	RondaActivity   RondaActivity         `gorm:"foreignKey:RondaActivityID"`
}

type RondaAttendances []RondaAttendance
