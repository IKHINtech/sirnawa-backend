package models

import "time"

type Notification struct {
	BaseModel
	UserID   string     `gorm:"type:uuid;not null" json:"user_id"`                  // Wajib, karena notifikasi harus punya target user
	HouseID  *string    `gorm:"type:uuid;null" json:"house_id"`                     // Opsional (bisa null)
	RtID     *string    `gorm:"type:uuid;null" json:"rt_id"`                        // Opsional (bisa null)
	Title    string     `gorm:"size:255;not null" json:"title"`                     // Judul notifikasi
	Body     string     `gorm:"type:text;not null" json:"body"`                     // Isi notifikasi
	Category string     `gorm:"size:50;not null;default:'general'" json:"category"` // "payment", "announcement", "event", etc.
	Data     JSONB      `gorm:"type:jsonb;null" json:"data"`                        // Payload tambahan (JSON)
	ReadAt   *time.Time `gorm:"index:idx_user_read_at;null" json:"read_at"`

	// Relasi (optional, jika butuh eager loading)
	User  User  `gorm:"foreignKey:UserID" json:"user,omitempty"`
	House House `gorm:"foreignKey:HouseID" json:"house,omitempty"`
	Rt    Rt    `gorm:"foreignKey:RtID" json:"rt,omitempty"`
}
