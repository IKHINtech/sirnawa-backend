package models

type User struct {
	BaseModel
	Email      string    `gorm:"uniqueIndex" json:"email"`
	Password   string    `gorm:"not null" json:"-"`
	Role       Role      `gorm:"type:role;not null" json:"role"`
	ResidentID *string   `gorm:"null" json:"resident_id"` // nullable
	Resident   *Resident `gorm:"foreignKey:ResidentID" json:"resident,omitempty"`
}

type Users []User
