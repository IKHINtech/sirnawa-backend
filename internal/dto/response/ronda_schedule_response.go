package response

import (
	"time"

	"github.com/IKHINtech/sirnawa-backend/internal/models"
)

type RondaScheduleResponse struct {
	BaseResponse
	Date    time.Time `json:"date"`
	GroupID string    `json:"group_id"`
}

type RondaScheduleResponses []RondaScheduleResponse

func RondaScheduleModelToRondaScheduleResponse(data *models.RondaSchedule) *RondaScheduleResponse {
	base := BaseResponse{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	return &RondaScheduleResponse{
		BaseResponse: base,
		Date:         data.Date,
		GroupID:      data.GroupID,
	}
}

func RondaScheduleListToResponse(data models.RondaSchedules) RondaScheduleResponses {
	var res RondaScheduleResponses
	for _, v := range data {
		res = append(res, *RondaScheduleModelToRondaScheduleResponse(&v))
	}
	return res
}
