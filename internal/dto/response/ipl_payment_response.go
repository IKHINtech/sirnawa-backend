package response

import (
	"time"

	"github.com/IKHINtech/sirnawa-backend/internal/models"
)

type IplPaymentResponse struct {
	BaseResponse
	IplBillID          string    `json:"ipl_bill_id"`
	PaidAt             time.Time `json:"paid_at,omitempty"`
	AmountPaid         int64     `json:"amount_paid"`
	IplPaymentMethod   string    `json:"payment_method"`
	Evidence           string    `json:"evidence"`
	VerifiedBy         *string   `json:"verified"`
	VerificationStatus string    `json:"verification_status"`
	Notes              string    `json:"notes"`
}

type IplPaymentResponses []IplPaymentResponse

func IplPaymentModelToIplPaymentResponse(data *models.IplPayment) *IplPaymentResponse {
	base := BaseResponse{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	return &IplPaymentResponse{
		BaseResponse:       base,
		IplBillID:          data.IplBillID,
		PaidAt:             data.PaidAt,
		AmountPaid:         data.AmountPaid,
		IplPaymentMethod:   data.IplPaymentMethod.ToString(),
		Evidence:           data.Evidence,
		Notes:              data.Notes,
		VerificationStatus: data.VerificationStatus.ToString(),
		VerifiedBy:         data.VerifiedBy,
	}
}

func IplPaymentListToResponse(data models.IplPayments) IplPaymentResponses {
	var res IplPaymentResponses
	for _, v := range data {
		res = append(res, *IplPaymentModelToIplPaymentResponse(&v))
	}
	return res
}
