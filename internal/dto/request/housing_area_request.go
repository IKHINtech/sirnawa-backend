package request

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type HousingAreaCreateRequest struct {
	Name string `json:"name"`
	RtID string `json:"rt_id"`
}

type HousingAreaUpdateRequset struct {
	ID string `json:"id"`
	HousingAreaCreateRequest
}

func HousingAreaUpdateRequsetToHousingAreaModel(data HousingAreaUpdateRequset) models.HousingArea {
	base := models.BaseModel{
		ID: data.ID,
	}
	return models.HousingArea{
		Name:      data.Name,
		BaseModel: base,
	}
}

func HousingAreaCreateRequestToHousingAreaModel(data HousingAreaCreateRequest) models.HousingArea {
	return models.HousingArea{
		Name: data.Name,
	}
}
