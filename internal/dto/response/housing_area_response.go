package response

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type HousingAreaResponse struct {
	BaseResponse
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
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
		Latitude:     data.Latitude,
		Longitude:    data.Longitude,
	}
}

func HousingAreaListToResponse(data models.HousingAreas) HousingAreaResponses {
	var res HousingAreaResponses
	for _, v := range data {
		res = append(res, *HousingAreaModelToHousingAreaResponse(&v))
	}
	return res
}
