package response

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type RwResponse struct {
	BaseResponse
	Name string `json:"name"`
}

type RwResponses []RwResponse

func RwModelToRwResponse(data *models.Rw) *RwResponse {
	base := BaseResponse{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	return &RwResponse{
		Name:         data.Name,
		BaseResponse: base,
	}
}

func RwListToResponse(data models.Rws) RwResponses {
	var res RwResponses
	for _, v := range data {
		res = append(res, *RwModelToRwResponse(&v))
	}
	return res
}
