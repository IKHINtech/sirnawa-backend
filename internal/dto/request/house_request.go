package request

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type HouseCreateRequest struct {
	BlockID string `json:"block_id"`
	Number  string `json:"number"`
	RtID    string `json:"rt_id"`
	RwID    string `json:"rw_id"`
	Status  string `json:"status"` // aktif / tidak aktif
}

type HouseUpdateRequset struct {
	ID string `json:"id"`
	HouseCreateRequest
}

func HouseUpdateRequsetToHouseModel(data HouseUpdateRequset) models.House {
	base := models.BaseModel{
		ID: data.ID,
	}
	return models.House{
		BaseModel: base,
		Number:    data.Number,
		RtID:      data.RtID,
		RwID:      data.RwID,
		Status:    models.HouseStatus(data.Status),
	}
}

func HouseCreateRequestToHouseModel(data HouseCreateRequest) models.House {
	return models.House{
		Number:  data.Number,
		Status:  models.HouseStatus(data.Status),
		BlockID: data.BlockID,
		RtID:    data.RtID,
		RwID:    data.RwID,
	}
}
