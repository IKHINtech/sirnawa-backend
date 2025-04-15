package response

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type HouseResponse struct {
	BaseResponse
	Number string `json:"number"`
	Status string `json:"status"`
}

type HouseResponses []HouseResponse

func HouseModelToHouseResponse(data *models.House) *HouseResponse {
	base := BaseResponse{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	return &HouseResponse{
		Number:       data.Number,
		Status:       string(data.Status),
		BaseResponse: base,
	}
}

func HouseListToResponse(data models.Houses) HouseResponses {
	var res HouseResponses
	for _, v := range data {
		res = append(res, *HouseModelToHouseResponse(&v))
	}
	return res
}
