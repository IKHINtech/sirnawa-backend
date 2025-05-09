package request

import (
	"time"

	"github.com/IKHINtech/sirnawa-backend/internal/models"
)

type IplRateCreateRequest struct {
	RtID      string    `json:"rt_id"`
	BlockID   *string   `json:"block_id"`
	HouseType *string   `json:"house_type"`
	Amount    int64     `json:"amount"`
	StartDate time.Time `json:"start_date"`
}

type IplRateUpdateRequset struct {
	ID string `json:"id"`
	IplRateCreateRequest
}

func IplRateUpdateRequsetToIplRateModel(data IplRateUpdateRequset) models.IplRate {
	base := models.BaseModel{
		ID: data.ID,
	}

	var houseType *models.HouseType
	if data.HouseType != nil {
		typeHouse := models.HouseType(*data.HouseType)
		houseType = &typeHouse
	}
	return models.IplRate{
		BaseModel: base,
		RtID:      data.RtID,
		BlockID:   data.BlockID,
		HouseType: houseType,
		Amount:    data.Amount,
		StartDate: data.StartDate,
	}
}

func IplRateCreateRequestToIplRateModel(data IplRateCreateRequest) models.IplRate {
	var houseType *models.HouseType
	if data.HouseType != nil {
		typeHouse := models.HouseType(*data.HouseType)
		houseType = &typeHouse
	}
	return models.IplRate{
		RtID:      data.RtID,
		BlockID:   data.BlockID,
		HouseType: houseType,
		Amount:    data.Amount,
		StartDate: data.StartDate,
	}
}
