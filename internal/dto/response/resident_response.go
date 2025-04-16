package response

import (
	"time"

	"github.com/IKHINtech/sirnawa-backend/internal/models"
)

type ResidentResponse struct {
	BaseResponse

	HouseID        string    `json:"house_id"`
	Name           string    `json:"name"`
	NIK            string    `json:"nik"`
	PhoneNumber    *string   `json:"phone_number"`
	BirthDate      time.Time `json:"birth_date"`
	Gender         string    `json:"gender"`
	Job            string    `json:"job"`
	IsHeadOfFamily bool      `json:"is_head_of_family"`
}

type ResidentResponses []ResidentResponse

func ResidentModelToResidentResponse(data *models.Resident) *ResidentResponse {
	base := BaseResponse{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	return &ResidentResponse{
		Name:           data.Name,
		NIK:            data.NIK,
		PhoneNumber:    data.PhoneNumber,
		BirthDate:      data.BirthDate,
		Gender:         data.Gender,
		Job:            data.Job,
		IsHeadOfFamily: data.IsHeadOfFamily,
		BaseResponse:   base,
	}
}

func ResidentListToResponse(data models.Residents) ResidentResponses {
	var res ResidentResponses
	for _, v := range data {
		res = append(res, *ResidentModelToResidentResponse(&v))
	}
	return res
}
