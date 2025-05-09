package response

import (
	"github.com/IKHINtech/sirnawa-backend/internal/models"
)

type RondaGroupMemberResponse struct {
	BaseResponse
	GroupID    string `json:"group_id"`
	ResidentID string `json:"resident_id"`
	HouseID    string `json:"house_id"`
}

type RondaGroupMemberDetailResponse struct {
	RondaGroupMemberResponse
	HouseResponse    `json:"house"`
	ResidentResponse `json:"resident"`
}

type RondaGroupMemberDetailRespones []RondaGroupMemberDetailResponse

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
		HouseID:      data.HouseID,
	}
}

func RondaGroupMemberListToResponse(data models.RondaGroupMembers) RondaGroupMemberResponses {
	var res RondaGroupMemberResponses
	for _, v := range data {
		res = append(res, *RondaGroupMemberModelToRondaGroupMemberResponse(&v))
	}
	return res
}

func MapRondaGroupMemberDetailResponse(data *models.RondaGroupMember) *RondaGroupMemberDetailResponse {
	if data == nil {
		return nil
	}
	return &RondaGroupMemberDetailResponse{
		RondaGroupMemberResponse: *RondaGroupMemberModelToRondaGroupMemberResponse(data),
		HouseResponse:            *HouseModelToHouseResponse(&data.House),
		ResidentResponse:         *ResidentModelToResidentResponse(&data.Resident),
	}
}
