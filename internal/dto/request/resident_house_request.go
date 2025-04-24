package request

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type ResidentHouseCreateRequest struct {
	ResidentID string `json:"resident_id"`
	HouseID    string `json:"house_id"`
	IsPrimary  bool   `json:"is_primary"` // Apakah ini rumah utama
}

type ResidentHouseUpdateRequset struct {
	ID string `json:"id"`
	ResidentHouseCreateRequest
}

func ResidentHouseUpdateRequsetToResidentHouseModel(data ResidentHouseUpdateRequset) models.ResidentHouse {
	base := models.BaseModel{
		ID: data.ID,
	}
	return models.ResidentHouse{
		BaseModel:  base,
		ResidentID: data.ResidentID,
		HouseID:    data.HouseID,
		IsPrimary:  data.IsPrimary,
	}
}

func ResidentHouseCreateRequestToResidentHouseModel(data ResidentHouseCreateRequest) models.ResidentHouse {
	return models.ResidentHouse{
		ResidentID: data.ResidentID,
		HouseID:    data.HouseID,
		IsPrimary:  data.IsPrimary,
	}
}
