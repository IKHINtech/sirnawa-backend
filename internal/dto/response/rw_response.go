package response

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type RwResponse struct {
	BaseResponse
	Name          string `json:"name"`
	HousingAreaID string `json:"housing_area_id"`
}

type RwResponses []RwResponse

func RwModelToRwResponse(data *models.Rw) *RwResponse {
	base := BaseResponse{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	return &RwResponse{
		Name:          data.Name,
		HousingAreaID: data.HousingAreaID,
		BaseResponse:  base,
	}
}

func RwListToResponse(data models.Rws) RwResponses {
	var res RwResponses
	for _, v := range data {
		res = append(res, *RwModelToRwResponse(&v))
	}
	return res
}
