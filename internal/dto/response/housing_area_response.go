package response

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type HousingAreaResponse struct {
	BaseResponse
	Name string `json:"name"`
}

type HousingAreaResponses []HousingAreaResponse

func HousingAreaModelToHousingAreaResponse(data *models.HousingArea) *HousingAreaResponse {
	base := BaseResponse{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	return &HousingAreaResponse{
		Name:         data.Name,
		BaseResponse: base,
	}
}

func HousingAreaListToResponse(data models.HousingAreas) HousingAreaResponses {
	var res HousingAreaResponses
	for _, v := range data {
		res = append(res, *HousingAreaModelToHousingAreaResponse(&v))
	}
	return res
}
