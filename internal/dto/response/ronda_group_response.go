package response

import (
	"github.com/IKHINtech/sirnawa-backend/internal/models"
)

type RondaGroupResponse struct {
	BaseResponse
	Name        string      `json:"name"`
	Order       uint        `json:"order"`
	RtID        string      `json:"rt_id"`
	Rt          *RtResponse `json:"rt"`
	TotalMember *int64      `json:"total_member"`
}

type RondaGroupDetailResponse struct {
	RondaGroupResponse
	Member RondaGroupMemberDetailRespones `json:"members"`
}

type RondaGroupResponses []RondaGroupResponse

func RondaGroupModelToRondaGroupResponse(data *models.RondaGroup, totalMember *int64) *RondaGroupResponse {
	if data == nil {
		return nil
	}
	base := BaseResponse{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	var rt *RtResponse
	if data.Rt.ID != "" {
		rt = RtModelToRtResponse(&data.Rt)
	}
	return &RondaGroupResponse{
		BaseResponse: base,
		Order:        data.Order,
		Name:         data.Name,
		RtID:         data.RtID,
		Rt:           rt,
		TotalMember:  totalMember,
	}
}

func RondaGroupListToResponse(data models.RondaGroups) RondaGroupResponses {
	var res RondaGroupResponses
	for _, v := range data {
		res = append(res, *RondaGroupModelToRondaGroupResponse(&v, nil))
	}
	return res
}

func MapRondaGroupDetailResponse(data *models.RondaGroup) *RondaGroupDetailResponse {
	res := &RondaGroupDetailResponse{
		RondaGroupResponse: *RondaGroupModelToRondaGroupResponse(data, nil),
	}

	members := make(RondaGroupMemberDetailRespones, len(data.Members))
	for i, member := range data.Members {
		members[i] = *MapRondaGroupMemberDetailResponse(&member)
	}

	res.Member = members
	return res
}
