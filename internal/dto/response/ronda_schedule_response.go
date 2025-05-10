package response

import (
	"time"

	"github.com/IKHINtech/sirnawa-backend/internal/models"
)

type RondaScheduleResponse struct {
	BaseResponse
	Date    time.Time           `json:"date"`
	GroupID string              `json:"group_id"`
	RtID    string              `json:"rt_id"`
	Group   *RondaGroupResponse `json:"group"`
	Rt      *RtResponse         `json:"rt"`
}

type RondaScheduleResponses []RondaScheduleResponse

func RondaScheduleModelToRondaScheduleResponse(data *models.RondaSchedule, totalMember *int64) *RondaScheduleResponse {
	if data == nil {
		return nil
	}
	base := BaseResponse{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	var group *RondaGroupResponse
	var rt *RtResponse

	if data.Group.ID != "" {
		x := RondaGroupModelToRondaGroupResponse(&data.Group, totalMember)
		group = x
	}

	if data.Rt.ID != "" {
		rt = RtModelToRtResponse(&data.Rt)
	}

	return &RondaScheduleResponse{
		BaseResponse: base,
		Date:         data.Date,
		GroupID:      data.GroupID,
		RtID:         data.RtID,
		Group:        group,
		Rt:           rt,
	}
}

func RondaScheduleListToResponse(data models.RondaSchedules) RondaScheduleResponses {
	var res RondaScheduleResponses
	for _, v := range data {
		res = append(res, *RondaScheduleModelToRondaScheduleResponse(&v, nil))
	}
	return res
}
