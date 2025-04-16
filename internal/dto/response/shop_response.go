package response

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type ShopResponse struct {
	BaseResponse
	UserID      uint   `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"` // aktif/nonaktif
}

type ShopResponses []ShopResponse

func ShopModelToShopResponse(data *models.Shop) *ShopResponse {
	base := BaseResponse{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	return &ShopResponse{
		Name:         data.Name,
		UserID:       data.UserID,
		Description:  data.Description,
		Status:       data.Status,
		BaseResponse: base,
	}
}

func ShopListToResponse(data models.Shops) ShopResponses {
	var res ShopResponses
	for _, v := range data {
		res = append(res, *ShopModelToShopResponse(&v))
	}
	return res
}
