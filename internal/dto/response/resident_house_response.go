package response

import (
	"github.com/IKHINtech/sirnawa-backend/internal/models"
)

type ResidentHouseResponse struct {
	BaseResponse
	ResidentID string `json:"resident_id"`
	HouseID    string `json:"house_id"`
	IsPrimary  bool   `json:"is_primary"` // Apakah ini rumah utama
}

type ResidentHouseDetailResponse struct {
	ResidentHouseResponse
	Resident ResidentResponse `json:"resident"`
}

type ResidentHouseFullResponse struct {
	ResidentHouseResponse
	House HouseResponseDetail `json:"house"`
}

type ResidentHouseResponses []ResidentHouseResponse

func ResidentHouseModelToResidentHouseResponse(data *models.ResidentHouse) *ResidentHouseResponse {
	base := BaseResponse{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	return &ResidentHouseResponse{
		BaseResponse: base,
		ResidentID:   data.ResidentID,
		HouseID:      data.HouseID,
		IsPrimary:    data.IsPrimary,
	}
}

func ResidentHouseListToResponse(data models.ResidentHouses) ResidentHouseResponses {
	var res ResidentHouseResponses
	for _, v := range data {
		res = append(res, *ResidentHouseModelToResidentHouseResponse(&v))
	}
	return res
}

func MapResidentHouseFullResponse(data models.ResidentHouse) ResidentHouseFullResponse {
	return ResidentHouseFullResponse{
		ResidentHouseResponse: *ResidentHouseModelToResidentHouseResponse(&data),
		House:                 *MapHouseDetailResponse(&data.House),
	}
}

func MapResidentHouseDetailResponse(data *models.ResidentHouse) *ResidentHouseDetailResponse {
	if data == nil {
		return nil
	}
	res := ResidentHouseModelToResidentHouseResponse(data)
	resident := ResidentModelToResidentResponse(&data.Resident)
	result := ResidentHouseDetailResponse{
		ResidentHouseResponse: *res,
		Resident:              *resident,
	}
	return &result
}
