package response

import (
	"time"

	"github.com/IKHINtech/sirnawa-backend/internal/models"
)

type BaseResponse struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BaseResponses []BaseResponse

func BaseModelToBaseResponse(data *models.BaseModel) *BaseResponse {
	if data != nil {
		return &BaseResponse{
			ID:        data.ID,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
		}
	}
	return nil
}

func BaseListToResponse(data models.BaseModels) *BaseResponses {
	var baseListResponse BaseResponses
	for _, item := range data {
		baseListResponse = append(baseListResponse, *BaseModelToBaseResponse(&item))
	}
	return &baseListResponse
}
