package models

type User struct {
	BaseModel
	Name       string    `json:"name"`
	Email      string    `gorm:"uniqueIndex" json:"email"`
	Password   string    `json:"-"`
	Role       string    `json:"role"`
	ResidentID *uint     `json:"resident_id"` // nullable
	Resident   *Resident `gorm:"foreignKey:ResidentID" json:"resident,omitempty"`
}
