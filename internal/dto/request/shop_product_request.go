package request

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type ShopProductCreateRequest struct {
	ShopID      uint    `json:"shop_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageURL    string  `json:"image_url"`
	Stock       int     `json:"stock"`
}

type ShopProductUpdateRequset struct {
	ID string `json:"id"`
	ShopProductCreateRequest
}

func ShopProductUpdateRequsetToShopProductModel(data ShopProductUpdateRequset) models.ShopProduct {
	base := models.BaseModel{
		ID: data.ID,
	}
	return models.ShopProduct{
		BaseModel:   base,
		ShopID:      data.ShopID,
		Name:        data.Name,
		Description: data.Description,
		Price:       data.Price,
		ImageURL:    data.ImageURL,
		Stock:       data.Stock,
	}
}

func ShopProductCreateRequestToShopProductModel(data ShopProductCreateRequest) models.ShopProduct {
	return models.ShopProduct{
		ShopID:      data.ShopID,
		Name:        data.Name,
		Description: data.Description,
		Price:       data.Price,
		ImageURL:    data.ImageURL,
		Stock:       data.Stock,
	}
}
