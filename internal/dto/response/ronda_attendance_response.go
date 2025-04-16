package response

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type RondaAttendanceResponse struct {
	BaseResponse
	RondaActivityID string `json:"ronda_activity_id"`
	ResidentID      string `json:"resident_id"`
	Status          string `json:"status"` // hadir / tidak_hadir
	Note            string `json:"note"`
}

type RondaAttendanceResponses []RondaAttendanceResponse

func RondaAttendanceModelToRondaAttendanceResponse(data *models.RondaAttendance) *RondaAttendanceResponse {
	base := BaseResponse{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	return &RondaAttendanceResponse{
		BaseResponse:    base,
		RondaActivityID: data.RondaActivityID,
		ResidentID:      data.ResidentID,
		Status:          string(data.Status),
		Note:            data.Note,
	}
}

func RondaAttendanceListToResponse(data models.RondaAttendances) RondaAttendanceResponses {
	var res RondaAttendanceResponses
	for _, v := range data {
		res = append(res, *RondaAttendanceModelToRondaAttendanceResponse(&v))
	}
	return res
}
