package response

import (
	"time"

	"github.com/IKHINtech/sirnawa-backend/internal/models"
)

type IplBillResponse struct {
	BaseResponse
	HouseID     string    `json:"house_id"`
	RtID        string    `json:"rt_id"`
	Month       int       `json:"month"`
	Year        int       `json:"year"`
	TotalAmount int64     `json:"total_amount"`
	AmountPaid  *int64    `json:"amount_paid"`
	BalanceDue  *int64    `json:"balance_due"`
	IplRateID   *string   `json:"ipl_rate_id"`
	Status      string    `json:"status"`
	DueDate     time.Time `json:"due_date"`
	Penalty     *string   `json:"penalty"`
}

type IplBillResponses []IplBillResponse

func IplBillModelToIplBillResponse(data *models.IplBill) *IplBillResponse {
	base := BaseResponse{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	return &IplBillResponse{
		BaseResponse: base,
		HouseID:      data.HouseID,
		RtID:         data.RtID,
		Month:        data.Month,
		Year:         data.Year,
		TotalAmount:  data.TotalAmount,
		AmountPaid:   data.AmountPaid,
		BalanceDue:   data.BalanceDue,
		IplRateID:    data.IplRateID,
		Status:       data.Status.ToString(),
		DueDate:      data.DueDate,
		Penalty:      data.Penalty,
	}
}

func IplBillListToResponse(data models.IplBills) IplBillResponses {
	var res IplBillResponses
	for _, v := range data {
		res = append(res, *IplBillModelToIplBillResponse(&v))
	}
	return res
}
