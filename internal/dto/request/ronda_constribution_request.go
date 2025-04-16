package request

import (
	"time"

	"github.com/IKHINtech/sirnawa-backend/internal/models"
)

type RondaConstributionCreateRequest struct {
	Name         string    `json:"name"`
	Date         time.Time `json:"date"`
	RondaGroupID string    `json:"ronda_group_id"`
	Total        float64   `json:"total"`
	TotalPenalty float64   `json:"total_penalty"`
}

type RondaConstributionUpdateRequset struct {
	ID string `json:"id"`
	RondaConstributionCreateRequest
}

func RondaConstributionUpdateRequsetToRondaConstributionModel(data RondaConstributionUpdateRequset) models.RondaConstribution {
	base := models.BaseModel{
		ID: data.ID,
	}
	return models.RondaConstribution{
		BaseModel: base,
		// TODO:
	}
}

func RondaConstributionCreateRequestToRondaConstributionModel(data RondaConstributionCreateRequest) models.RondaConstribution {
	return models.RondaConstribution{}
}
