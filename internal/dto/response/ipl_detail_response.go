package response

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type IplDetailResponse struct {
	BaseResponse
	IplID     string  `json:"ipl_id"`
	Note      string  `json:"note"`
	SubAmount float64 `json:"sub_amount"`
}

type IplDetailResponses []IplDetailResponse

func IplDetailModelToIplDetailResponse(data *models.IplDetail) *IplDetailResponse {
	base := BaseResponse{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	return &IplDetailResponse{
		BaseResponse: base,
		IplID:        data.IplID,
		Note:         data.Note,
		SubAmount:    data.SubAmount,
	}
}

func IplDetailListToResponse(data models.IplDetails) IplDetailResponses {
	var res IplDetailResponses
	for _, v := range data {
		res = append(res, *IplDetailModelToIplDetailResponse(&v))
	}
	return res
}
