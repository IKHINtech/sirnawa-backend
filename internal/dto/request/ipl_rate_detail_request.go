package request

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type IplRateDetailCreateRequest struct {
	IplRateID string `json:"ipl_rate_id"`
	ItemID    string `json:"item_id"`
	Amount    int64  `json:"amount"`
}

type IplRateDetailUpdateRequset struct {
	ID string `json:"id"`
	IplRateDetailCreateRequest
}

func IplRateDetailUpdateRequsetToIplRateDetailModel(data IplRateDetailUpdateRequset) models.IplRateDetail {
	base := models.BaseModel{
		ID: data.ID,
	}
	return models.IplRateDetail{
		BaseModel: base,
		IplRateID: data.IplRateID,
		ItemID:    data.ItemID,
		Amount:    data.Amount,
	}
}

func IplRateDetailCreateRequestToIplRateDetailModel(data IplRateDetailCreateRequest) models.IplRateDetail {
	return models.IplRateDetail{
		IplRateID: data.IplRateID,
		ItemID:    data.ItemID,
		Amount:    data.Amount,
	}
}
