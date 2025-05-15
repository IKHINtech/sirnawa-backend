package request

import (
	"time"

	"github.com/IKHINtech/sirnawa-backend/internal/models"
)

type IplBillCreateRequest struct {
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

type IplBillUpdateRequset struct {
	ID string `json:"id"`
	IplBillCreateRequest
}

type IplBillGenerator struct {
	RtID       string   `json:"rt_id"`
	Month      int      `json:"month"`
	Year       int      `json:"year"`
	IplRateID  string   `json:"ipl_rate_id"`
	HouseIDs   []string `json:"house_ids"`
	IsAllHouse bool     `json:"is_all_house"`
	DueDate    int      `json:"due_date"`
}

func IplBillUpdateRequsetToIplBillModel(data IplBillUpdateRequset) models.IplBill {
	base := models.BaseModel{
		ID: data.ID,
	}
	return models.IplBill{
		BaseModel:   base,
		HouseID:     data.HouseID,
		RtID:        data.RtID,
		Month:       data.Month,
		Year:        data.Year,
		TotalAmount: data.TotalAmount,
		AmountPaid:  data.AmountPaid,
		BalanceDue:  data.BalanceDue,
		IplRateID:   data.IplRateID,
		Status:      models.IplBillStatus(data.Status),
		DueDate:     data.DueDate,
		Penalty:     data.Penalty,
	}
}

func IplBillCreateRequestToIplBillModel(data IplBillCreateRequest) models.IplBill {
	return models.IplBill{
		HouseID:     data.HouseID,
		RtID:        data.RtID,
		Month:       data.Month,
		Year:        data.Year,
		TotalAmount: data.TotalAmount,
		AmountPaid:  data.AmountPaid,
		BalanceDue:  data.BalanceDue,
		IplRateID:   data.IplRateID,
		Status:      models.IplBillStatus(data.Status),
		DueDate:     data.DueDate,
		Penalty:     data.Penalty,
	}
}
