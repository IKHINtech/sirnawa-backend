package request

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type BlockCreateRequest struct {
	Name string `json:"name"`
	RtID string `json:"rt_id"`
}

type BlockUpdateRequset struct {
	ID string `json:"id"`
	BlockCreateRequest
}

func BlockUpdateRequsetToBlockModel(data BlockUpdateRequset) models.Block {
	base := models.BaseModel{
		ID: data.ID,
	}
	return models.Block{
		Name:      data.Name,
		RtID:      data.RtID,
		BaseModel: base,
	}
}

func BlockCreateRequestToBlockModel(data BlockCreateRequest) models.Block {
	return models.Block{
		Name: data.Name,
		RtID: data.RtID,
	}
}
