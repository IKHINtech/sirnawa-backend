package request

import (
	"time"

	"github.com/IKHINtech/sirnawa-backend/internal/models"
)

type IplPaymentCreateRequest struct {
	HouseID string     `json:"house_id"`
	Month   int        `json:"month"`
	Year    int        `json:"year"`
	Amount  float64    `json:"amount"`
	PaidAt  *time.Time `json:"paid_at,omitempty"`
}

type IplPaymentUpdateRequset struct {
	ID string `json:"id"`
	IplPaymentCreateRequest
}

func IplPaymentUpdateRequsetToIplPaymentModel(data IplPaymentUpdateRequset) models.IplPayment {
	base := models.BaseModel{
		ID: data.ID,
	}
	return models.IplPayment{
		HouseID:   data.HouseID,
		Month:     data.Month,
		Year:      data.Year,
		Amount:    data.Amount,
		BaseModel: base,
	}
}

func IplPaymentCreateRequestToIplPaymentModel(data IplPaymentCreateRequest) models.IplPayment {
	return models.IplPayment{
		HouseID: data.HouseID,
		Month:   data.Month,
		Year:    data.Year,
		Amount:  data.Amount,
	}
}
