package request

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type IplCreateRequest struct {
	RtID     string                   `json:"rt_id"`
	Amount   float64                  `json:"amount"`
	IsActive bool                     `json:"is_active"`
	Note     string                   `json:"note"`
	Items    []IplDetailCreateRequest `json:"items"`
}

type IplUpdateRequset struct {
	ID       string                   `json:"id"`
	RtID     string                   `json:"rt_id"`
	Amount   float64                  `json:"amount"`
	IsActive bool                     `json:"is_active"`
	Note     string                   `json:"note"`
	Items    []IplDetailCreateRequest `json:"items"`
}

func IplUpdateRequsetToIplModel(data IplUpdateRequset) models.Ipl {
	base := models.BaseModel{
		ID: data.ID,
	}
	return models.Ipl{
		BaseModel: base,
		RtID:      data.RtID,
		Amount:    data.Amount,
		IsActive:  data.IsActive,
		Note:      data.Note,
	}
}

func IplCreateRequestToIplModel(data IplCreateRequest) models.Ipl {
	return models.Ipl{
		RtID:     data.RtID,
		Amount:   data.Amount,
		IsActive: data.IsActive,
		Note:     data.Note,
	}
}
