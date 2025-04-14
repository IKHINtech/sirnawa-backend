package request

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type BaseCreateRequest struct{}

type BaseUpdateRequset struct {
	ID string `json:"id"`
	BaseCreateRequest
}

func BaseCreateRequestToBaseModel(data BaseCreateRequest) *models.BaseModel {
	return &models.BaseModel{}
}

func BaseUpdateRequsetToBaseModel(data BaseUpdateRequset) *models.BaseModel {
	return &models.BaseModel{}
}
