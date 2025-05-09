package response

import (
	"github.com/IKHINtech/sirnawa-backend/internal/models"
)

type RondaGroupResponse struct {
	BaseResponse
	Name  string `json:"name"`
	Order uint   `json:"order"`
	RtID  string `json:"rt_id"`
}

type RondaGroupDetailResponse struct {
	RondaGroupResponse
	Member RondaGroupMemberDetailRespones `json:"members"`
}

type RondaGroupResponses []RondaGroupResponse

func RondaGroupModelToRondaGroupResponse(data *models.RondaGroup) *RondaGroupResponse {
	base := BaseResponse{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	return &RondaGroupResponse{
		BaseResponse: base,
		Order:        data.Order,
		Name:         data.Name,
		RtID:         data.RtID,
	}
}

func RondaGroupListToResponse(data models.RondaGroups) RondaGroupResponses {
	var res RondaGroupResponses
	for _, v := range data {
		res = append(res, *RondaGroupModelToRondaGroupResponse(&v))
	}
	return res
}

func MapRondaGroupDetailResponse(data *models.RondaGroup) *RondaGroupDetailResponse {
	res := &RondaGroupDetailResponse{
		RondaGroupResponse: *RondaGroupModelToRondaGroupResponse(data),
	}

	members := make(RondaGroupMemberDetailRespones, len(data.Members))
	for i, member := range data.Members {
		members[i] = *MapRondaGroupMemberDetailResponse(&member)
	}

	res.Member = members
	return res
}
