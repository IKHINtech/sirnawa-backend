package response

import (
	"time"

	"github.com/IKHINtech/sirnawa-backend/internal/models"
)

type RondaConstributionResponse struct {
	BaseResponse
	Date         time.Time `json:"date"`
	RondaGroupID string    `json:"ronda_group_id"`
	Total        float64   `json:"total"`
	TotalPenalty float64   `json:"total_penalty"`
}

type RondaConstributionResponses []RondaConstributionResponse

func RondaConstributionModelToRondaConstributionResponse(data *models.RondaConstribution) *RondaConstributionResponse {
	base := BaseResponse{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	return &RondaConstributionResponse{
		BaseResponse: base,
		Date:         data.Date,
		RondaGroupID: data.RondaGroupID,
		Total:        data.Total,
		TotalPenalty: data.TotalPenalty,
	}
}

func RondaConstributionListToResponse(data models.RondaConstributions) RondaConstributionResponses {
	var res RondaConstributionResponses
	for _, v := range data {
		res = append(res, *RondaConstributionModelToRondaConstributionResponse(&v))
	}
	return res
}
