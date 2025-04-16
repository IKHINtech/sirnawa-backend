package request

import (
	"time"

	"github.com/IKHINtech/sirnawa-backend/internal/models"
)

type RondaActivityCreateRequest struct {
	RondaGroupID string    `json:"ronda_group_id"`
	Date         time.Time `json:"date"`
	Description  string    `json:"description"`
	CreatedBy    string    `json:"created_by"`
}

type RondaActivityUpdateRequset struct {
	ID string `json:"id"`
	RondaActivityCreateRequest
}

func RondaActivityUpdateRequsetToRondaActivityModel(data RondaActivityUpdateRequset) models.RondaActivity {
	base := models.BaseModel{
		ID: data.ID,
	}
	return models.RondaActivity{
		BaseModel:    base,
		RondaGroupID: data.RondaGroupID,
		Date:         data.Date,
		Description:  data.Description,
		CreatedBy:    data.CreatedBy,
	}
}

func RondaActivityCreateRequestToRondaActivityModel(data RondaActivityCreateRequest) models.RondaActivity {
	return models.RondaActivity{
		RondaGroupID: data.RondaGroupID,
		Date:         data.Date,
		Description:  data.Description,
		CreatedBy:    data.CreatedBy,
	}
}
