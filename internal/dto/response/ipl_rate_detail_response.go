package response

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type IplRateDetailResponse struct {
	BaseResponse
	IplRateID string `json:"ipl_rate_id"`
	ItemID    string `json:"item_id"`
	Amount    int64  `json:"amount"`
}

type IplRateDetailResponses []IplRateDetailResponse

func IplRateDetailModelToIplRateDetailResponse(data *models.IplRateDetail) *IplRateDetailResponse {
	base := BaseResponse{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	return &IplRateDetailResponse{
		BaseResponse: base,
		IplRateID:    data.IplRateID,
		ItemID:       data.ItemID,
		Amount:       data.Amount,
	}
}

func IplRateDetailListToResponse(data models.IplRateDetails) IplRateDetailResponses {
	var res IplRateDetailResponses
	for _, v := range data {
		res = append(res, *IplRateDetailModelToIplRateDetailResponse(&v))
	}
	return res
}
