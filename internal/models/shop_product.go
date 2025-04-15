package models

type ShopProduct struct {
	BaseModel
	ShopID      uint    `json:"shop_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageURL    string  `json:"image_url"`
	Stock       int     `json:"stock"`
}

type ShopProducts []ShopProduct
