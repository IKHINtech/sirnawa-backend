package request

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type IplDetailCreateRequest struct {
	IplID     string  `json:"ipl_id"`
	Note      string  `json:"note"`
	SubAmount float64 `json:"sub_amount"`
}

type IplDetailUpdateRequset struct {
	ID string `json:"id"`
	IplDetailCreateRequest
}

func IplDetailUpdateRequsetToIplDetailModel(data IplDetailUpdateRequset) models.IplDetail {
	base := models.BaseModel{
		ID: data.ID,
	}
	return models.IplDetail{
		BaseModel: base,
		IplID:     data.IplID,
		Note:      data.Note,
		SubAmount: data.SubAmount,
	}
}

func IplDetailCreateRequestToIplDetailModel(data IplDetailCreateRequest) models.IplDetail {
	return models.IplDetail{
		IplID:     data.IplID,
		Note:      data.Note,
		SubAmount: data.SubAmount,
	}
}
