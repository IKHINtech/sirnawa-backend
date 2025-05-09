package request

import (
	"time"

	"github.com/IKHINtech/sirnawa-backend/internal/models"
)

type IplPaymentCreateRequest struct {
	IplBillID        string    `json:"ipl_bill_id" `
	PaidAt           time.Time `json:"paid_at"`
	AmountPaid       int64     `json:"amount_paid" `
	IplPaymentMethod string    `json:"payment_method" `
	Evidence         string    `json:"evidence"`
	Notes            string    `json:"notes"`
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
		BaseModel:        base,
		IplBillID:        data.IplBillID,
		PaidAt:           data.PaidAt,
		AmountPaid:       data.AmountPaid,
		IplPaymentMethod: models.PaymentMethod(data.IplPaymentMethod),
		Evidence:         data.Evidence,
		Notes:            data.Notes,
	}
}

func IplPaymentCreateRequestToIplPaymentModel(data IplPaymentCreateRequest) models.IplPayment {
	return models.IplPayment{
		IplBillID:        data.IplBillID,
		PaidAt:           data.PaidAt,
		AmountPaid:       data.AmountPaid,
		IplPaymentMethod: models.PaymentMethod(data.IplPaymentMethod),
		Evidence:         data.Evidence,
		Notes:            data.Notes,
	}
}
