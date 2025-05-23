package request

import (
	"time"

	"github.com/IKHINtech/sirnawa-backend/internal/models"
)

type RondaScheduleCreateRequest struct {
	Date    time.Time `json:"date"`
	GroupID string    `json:"group_id"`
	RtID    string    `json:"rt_id"`
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
		BaseModel: base,
		Date:      data.Date,
		GroupID:   data.GroupID,
		RtID:      data.RtID,
	}
}

func RondaScheduleCreateRequestToRondaScheduleModel(data RondaScheduleCreateRequest) models.RondaSchedule {
	return models.RondaSchedule{
		Date:    data.Date,
		GroupID: data.GroupID,
		RtID:    data.RtID,
	}
}
