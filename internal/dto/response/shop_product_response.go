package response

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type ShopProductResponse struct {
	BaseResponse
	ShopID      uint    `json:"shop_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageURL    string  `json:"image_url"`
	Stock       int     `json:"stock"`
}

type ShopProductResponses []ShopProductResponse

func ShopProductModelToShopProductResponse(data *models.ShopProduct) *ShopProductResponse {
	base := BaseResponse{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	return &ShopProductResponse{
		BaseResponse: base,
		Name:         data.Name,
		Description:  data.Description,
		ShopID:       data.ShopID,
		Price:        data.Price,
		ImageURL:     data.ImageURL,
		Stock:        data.Stock,
	}
}

func ShopProductListToResponse(data models.ShopProducts) ShopProductResponses {
	var res ShopProductResponses
	for _, v := range data {
		res = append(res, *ShopProductModelToShopProductResponse(&v))
	}
	return res
}
