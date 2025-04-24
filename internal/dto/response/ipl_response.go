package response

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type IplResponse struct {
	BaseResponse
	RtID     string  `json:"rt_id"`
	Amount   float64 `json:"amount"`
	IsActive bool    `json:"is_active"`
	Note     string  `json:"note"`
}

type IplFullResponse struct {
	IplResponse
	Items IplDetailResponses `json:"items"`
}

type IplResponses []IplResponse

func IplModelToIplResponse(data *models.Ipl) *IplResponse {
	base := BaseResponse{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	return &IplResponse{
		BaseResponse: base,
		RtID:         data.RtID,
		Amount:       data.Amount,
		IsActive:     data.IsActive,
		Note:         data.Note,
	}
}

func IplListToResponse(data models.Ipls) IplResponses {
	var res IplResponses
	for _, v := range data {
		res = append(res, *IplModelToIplResponse(&v))
	}
	return res
}

func IplModelToIplFullResponse(data *models.Ipl) *IplFullResponse {
	if data == nil {
		return nil
	}

	ipl := IplModelToIplResponse(data)
	items := IplDetailListToResponse(data.Items)
	return &IplFullResponse{
		IplResponse: *ipl,
		Items:       items,
	}
}
