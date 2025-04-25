package request

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type ShopCreateRequest struct {
	UserID      uint   `json:"user_id"`
	RtID        string `json:"rt_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"` // aktif/nonaktif
}

type ShopUpdateRequset struct {
	ID string `json:"id"`
	ShopCreateRequest
}

func ShopUpdateRequsetToShopModel(data ShopUpdateRequset) models.Shop {
	base := models.BaseModel{
		ID: data.ID,
	}
	return models.Shop{
		Name:        data.Name,
		BaseModel:   base,
		UserID:      data.UserID,
		RtID:        data.RtID,
		Description: data.Description,
		Status:      data.Status,
	}
}

func ShopCreateRequestToShopModel(data ShopCreateRequest) models.Shop {
	return models.Shop{
		Name:        data.Name,
		RtID:        data.RtID,
		UserID:      data.UserID,
		Description: data.Description,
		Status:      data.Status,
	}
}
