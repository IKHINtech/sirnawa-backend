package response

import (
	"time"

	"github.com/IKHINtech/sirnawa-backend/internal/models"
)

type IplRateResponse struct {
	BaseResponse
	RtID      string    `json:"rt_id"`
	BlockID   *string   `json:"block_id"`
	HouseType *string   `json:"house_type"`
	Amount    int64     `json:"amount"`
	StartDate time.Time `json:"start_date"`
}

type IplRateResponses []IplRateResponse

func IplRateModelToIplRateResponse(data *models.IplRate) *IplRateResponse {
	base := BaseResponse{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	var houseType *string

	if data.HouseType != nil {
		v := models.HouseType(*data.HouseType).ToString()
		houseType = &v
	}
	return &IplRateResponse{
		BaseResponse: base,
		RtID:         data.RtID,
		BlockID:      data.BlockID,
		HouseType:    houseType,
		Amount:       data.Amount,
		StartDate:    data.StartDate,
	}
}

func IplRateListToResponse(data models.IplRates) IplRateResponses {
	var res IplRateResponses
	for _, v := range data {
		res = append(res, *IplRateModelToIplRateResponse(&v))
	}
	return res
}
