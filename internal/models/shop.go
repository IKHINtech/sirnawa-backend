package models

type Shop struct {
	BaseModel
	UserID      uint          `json:"user_id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Status      string        `json:"status"` // aktif/nonaktif
	Products    []ShopProduct `gorm:"foreignKey:ShopID" json:"products,omitempty"`
}
type Shops []Shop
