package request

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type RondaGroupCreateRequest struct {
	Name string `json:"name"`
	RtID string `json:"rt_id"`
}

type RondaGroupUpdateRequset struct {
	ID string `json:"id"`
	RondaGroupCreateRequest
}

func RondaGroupUpdateRequsetToRondaGroupModel(data RondaGroupUpdateRequset) models.RondaGroup {
	base := models.BaseModel{
		ID: data.ID,
	}
	return models.RondaGroup{
		Name:      data.Name,
		RtID:      data.RtID,
		BaseModel: base,
	}
}

func RondaGroupCreateRequestToRondaGroupModel(data RondaGroupCreateRequest) models.RondaGroup {
	return models.RondaGroup{
		Name: data.Name,
		RtID: data.RtID,
	}
}
