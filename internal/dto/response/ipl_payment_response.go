package response

import (
	"time"

	"github.com/IKHINtech/sirnawa-backend/internal/models"
)

type IplPaymentResponse struct {
	BaseResponse

	HouseID string     `json:"house_id"`
	Month   int        `json:"month"`
	Year    int        `json:"year"`
	Amount  float64    `json:"amount"`
	Status  string     `json:"status"` // paid/unpaid
	PaidAt  *time.Time `json:"paid_at,omitempty"`
}

type IplPaymentResponses []IplPaymentResponse

func IplPaymentModelToIplPaymentResponse(data *models.IplPayment) *IplPaymentResponse {
	base := BaseResponse{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	return &IplPaymentResponse{
		BaseResponse: base,
		HouseID:      data.HouseID,
		Month:        data.Month,
		Year:         data.Year,
		Amount:       data.Amount,
		Status:       string(data.Status),
		PaidAt:       data.PaidAt,
	}
}

func IplPaymentListToResponse(data models.IplPayments) IplPaymentResponses {
	var res IplPaymentResponses
	for _, v := range data {
		res = append(res, *IplPaymentModelToIplPaymentResponse(&v))
	}
	return res
}
