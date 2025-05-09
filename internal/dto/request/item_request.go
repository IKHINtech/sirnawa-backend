package request

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type ItemCreateRequest struct {
	Name        string `json:"name"`
	RtID        string `json:"rt_id"`
	Description string `json:"description"`
	ItemmType   string `json:"item_type"`
}

type ItemUpdateRequset struct {
	ID string `json:"id"`
	ItemCreateRequest
}

func ItemUpdateRequsetToItemModel(data ItemUpdateRequset) models.Item {
	base := models.BaseModel{
		ID: data.ID,
	}
	return models.Item{
		Name:        data.Name,
		RtID:        data.RtID,
		ItemType:    models.ItemType(data.ItemmType),
		Description: data.Description,
		BaseModel:   base,
	}
}

func ItemCreateRequestToItemModel(data ItemCreateRequest) models.Item {
	return models.Item{
		Name:        data.Name,
		RtID:        data.RtID,
		ItemType:    models.ItemType(data.ItemmType),
		Description: data.Description,
	}
}
