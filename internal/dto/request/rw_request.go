package request

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type RwCreateRequest struct {
	Name          string `json:"name"`
	HousingAreaID string `json:"housing_area_id"`
}

type RwUpdateRequset struct {
	ID string `json:"id"`
	RwCreateRequest
}

func RwUpdateRequsetToRwModel(data RwUpdateRequset) models.Rw {
	base := models.BaseModel{
		ID: data.ID,
	}
	return models.Rw{
		Name:          data.Name,
		HousingAreaID: data.HousingAreaID,
		BaseModel:     base,
	}
}

func RwCreateRequestToRwModel(data RwCreateRequest) models.Rw {
	return models.Rw{
		Name:          data.Name,
		HousingAreaID: data.HousingAreaID,
	}
}
