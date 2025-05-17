package models

type User struct {
	BaseModel
	Email      *string   `gorm:"uniqueIndex" json:"email"`
	IsActive   bool      `gorm:"default:false"`
	Password   string    `gorm:"not null" json:"-"`
	Role       Role      `gorm:"type:role;not null;default:warga" json:"role"`
	PhotoUrl   *string   `gorm:"null" json:"photo_url"`
	ResidentID *string   `gorm:"null" json:"resident_id"` // nullable
	Resident   *Resident `gorm:"foreignKey:ResidentID" json:"resident,omitempty"`
	UserRTs    UserRTs
}

type Users []User
