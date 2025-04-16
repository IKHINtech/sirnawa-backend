package request

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type RondaAttendanceCreateRequest struct {
	RondaActivityID string `json:"ronda_activity_id"`
	ResidentID      string `json:"resident_id"`
	Status          string `json:"status"` // hadir / tidak_hadir
	Note            string `json:"note"`
}

type RondaAttendanceUpdateRequset struct {
	ID string `json:"id"`
	RondaAttendanceCreateRequest
}

func RondaAttendanceUpdateRequsetToRondaAttendanceModel(data RondaAttendanceUpdateRequset) models.RondaAttendance {
	base := models.BaseModel{
		ID: data.ID,
	}
	return models.RondaAttendance{
		BaseModel:       base,
		RondaActivityID: data.RondaActivityID,
		ResidentID:      data.ResidentID,
		Status:          models.RondaAttendanceStatus(data.Status),
		Note:            data.Note,
	}
}

func RondaAttendanceCreateRequestToRondaAttendanceModel(data RondaAttendanceCreateRequest) models.RondaAttendance {
	return models.RondaAttendance{
		RondaActivityID: data.RondaActivityID,
		ResidentID:      data.ResidentID,
		Status:          models.RondaAttendanceStatus(data.Status),
		Note:            data.Note,
	}
}
