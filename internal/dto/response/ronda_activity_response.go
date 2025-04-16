package response

import (
	"time"

	"github.com/IKHINtech/sirnawa-backend/internal/models"
)

type RondaActivityResponse struct {
	BaseResponse
	RondaGroupID string    `json:"ronda_group_id"`
	Date         time.Time `gorm:"not null" json:"date"`
	Description  string    `gorm:"type:text" json:"description"`
	CreatedBy    string    `gorm:"not null" json:"created_by"`
}

type RondaActivityResponses []RondaActivityResponse

func RondaActivityModelToRondaActivityResponse(data *models.RondaActivity) *RondaActivityResponse {
	base := BaseResponse{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	return &RondaActivityResponse{
		BaseResponse: base,
		RondaGroupID: data.RondaGroupID,
		Date:         data.Date,
		Description:  data.Description,
		CreatedBy:    data.CreatedBy,
	}
}

func RondaActivityListToResponse(data models.RondaActivitys) RondaActivityResponses {
	var res RondaActivityResponses
	for _, v := range data {
		res = append(res, *RondaActivityModelToRondaActivityResponse(&v))
	}
	return res
}
