package request

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type RtCreateRequest struct {
	Name string `json:"name"`
	RwID string `json:"rw_id"`
}

type RtUpdateRequset struct {
	ID string `json:"id"`
	RtCreateRequest
}

func RtUpdateRequsetToRtModel(data RtUpdateRequset) models.Rt {
	base := models.BaseModel{
		ID: data.ID,
	}
	return models.Rt{
		Name:      data.Name,
		RwID:      data.RwID,
		BaseModel: base,
	}
}

func RtCreateRequestToRtModel(data RtCreateRequest) models.Rt {
	return models.Rt{
		Name: data.Name,
		RwID: data.RwID,
	}
}
