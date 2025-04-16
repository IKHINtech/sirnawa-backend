package response

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type RondaGroupMemberResponse struct {
	BaseResponse
	GroupID    string `json:"group_id"`
	ResidentID string `json:"resident_id"`
}

type RondaGroupMemberResponses []RondaGroupMemberResponse

func RondaGroupMemberModelToRondaGroupMemberResponse(data *models.RondaGroupMember) *RondaGroupMemberResponse {
	base := BaseResponse{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	return &RondaGroupMemberResponse{
		BaseResponse: base,
		GroupID:      data.GroupID,
		ResidentID:   data.ResidentID,
	}
}

func RondaGroupMemberListToResponse(data models.RondaGroupMembers) RondaGroupMemberResponses {
	var res RondaGroupMemberResponses
	for _, v := range data {
		res = append(res, *RondaGroupMemberModelToRondaGroupMemberResponse(&v))
	}
	return res
}
