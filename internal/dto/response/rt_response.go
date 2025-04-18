package response

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type RtResponse struct {
	BaseResponse
	Name string `json:"name"`
	RwID string `json:"rw_id"`
}

type RtResponses []RtResponse

func RtModelToRtResponse(data *models.Rt) *RtResponse {
	base := BaseResponse{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	return &RtResponse{
		Name:         data.Name,
		BaseResponse: base,
		RwID:         data.RwID,
	}
}

func RtListToResponse(data models.Rts) RtResponses {
	var res RtResponses
	for _, v := range data {
		res = append(res, *RtModelToRtResponse(&v))
	}
	return res
}
