package response

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type RondaGroupResponse struct {
	BaseResponse
	Name  string `json:"name"`
	Order uint   `json:"order"`
}

type RondaGroupResponses []RondaGroupResponse

func RondaGroupModelToRondaGroupResponse(data *models.RondaGroup) *RondaGroupResponse {
	base := BaseResponse{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	return &RondaGroupResponse{
		Name:         data.Name,
		BaseResponse: base,
		Order:        data.Order,
	}
}

func RondaGroupListToResponse(data models.RondaGroups) RondaGroupResponses {
	var res RondaGroupResponses
	for _, v := range data {
		res = append(res, *RondaGroupModelToRondaGroupResponse(&v))
	}
	return res
}
