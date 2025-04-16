package request

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type RondaGroupMemberCreateRequest struct {
	GroupID    string `json:"group_id"`
	ResidentID string `json:"resident_id"`
}

type RondaGroupMemberUpdateRequset struct {
	ID string `json:"id"`
	RondaGroupMemberCreateRequest
}

func RondaGroupMemberUpdateRequsetToRondaGroupMemberModel(data RondaGroupMemberUpdateRequset) models.RondaGroupMember {
	base := models.BaseModel{
		ID: data.ID,
	}
	return models.RondaGroupMember{
		BaseModel:  base,
		GroupID:    data.GroupID,
		ResidentID: data.ResidentID,
	}
}

func RondaGroupMemberCreateRequestToRondaGroupMemberModel(data RondaGroupMemberCreateRequest) models.RondaGroupMember {
	return models.RondaGroupMember{
		GroupID:    data.GroupID,
		ResidentID: data.ResidentID,
	}
}
