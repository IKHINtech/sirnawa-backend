package request

import (
	"time"

	"github.com/IKHINtech/sirnawa-backend/internal/models"
)

type RondaScheduleCreateRequest struct {
	EfektifDate time.Time `json:"efektif_date"`
	GroupID     string    `json:"group_id"`
}

type RondaScheduleUpdateRequset struct {
	ID string `json:"id"`
	RondaScheduleCreateRequest
}

func RondaScheduleUpdateRequsetToRondaScheduleModel(data RondaScheduleUpdateRequset) models.RondaSchedule {
	base := models.BaseModel{
		ID: data.ID,
	}
	return models.RondaSchedule{
		BaseModel:   base,
		EfektifDate: data.EfektifDate,
		GroupID:     data.GroupID,
	}
}

func RondaScheduleCreateRequestToRondaScheduleModel(data RondaScheduleCreateRequest) models.RondaSchedule {
	return models.RondaSchedule{
		EfektifDate: data.EfektifDate,
		GroupID:     data.GroupID,
	}
}
