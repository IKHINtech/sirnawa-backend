package request

import (
	"time"

	"github.com/IKHINtech/sirnawa-backend/internal/models"
)

type ResidentCreateRequest struct {
	HouseID        string    `json:"house_id"`
	Name           string    `json:"name"`
	Email          *string   `json:"email"`
	Role           string    `json:"role" example:"warga;admin_rt;sekretaris;bendahara;wakil_rt"`
	NIK            string    `json:"nik"`
	PhoneNumber    *string   `json:"phone_number"`
	BirthDate      time.Time `json:"birth_date"`
	Gender         string    `json:"gender"`
	Job            string    `json:"job"`
	IsHeadOfFamily bool      `json:"is_head_of_family"`
}

type ResidentUpdateRequset struct {
	ID string `json:"id"`
	ResidentCreateRequest
}

func ResidentUpdateRequsetToResidentModel(data ResidentUpdateRequset) models.Resident {
	base := models.BaseModel{
		ID: data.ID,
	}
	return models.Resident{
		Name:           data.Name,
		NIK:            data.NIK,
		PhoneNumber:    data.PhoneNumber,
		BirthDate:      data.BirthDate,
		Gender:         data.Gender,
		HouseID:        data.HouseID,
		Job:            data.Job,
		IsHeadOfFamily: data.IsHeadOfFamily,
		BaseModel:      base,
	}
}

func ResidentCreateRequestToResidentModel(data ResidentCreateRequest) models.Resident {
	return models.Resident{
		Name:           data.Name,
		NIK:            data.NIK,
		PhoneNumber:    data.PhoneNumber,
		BirthDate:      data.BirthDate,
		Gender:         data.Gender,
		HouseID:        data.HouseID,
		Job:            data.Job,
		IsHeadOfFamily: data.IsHeadOfFamily,
	}
}
