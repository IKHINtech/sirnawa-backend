package request

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type IplBillDetailCreateRequest struct {
	IplBillID string `json:"ipl_bill_id"`
	ItemID    string `json:"item_id"`
	Note      string `json:"note"`
	SubAmount int64  `json:"sub_amount"`
}

type IplBillDetailUpdateRequset struct {
	ID string `json:"id"`
	IplBillDetailCreateRequest
}

func IplBillDetailUpdateRequsetToIplBillDetailModel(data IplBillDetailUpdateRequset) models.IplBillDetail {
	base := models.BaseModel{
		ID: data.ID,
	}
	return models.IplBillDetail{
		BaseModel: base,
		IplBillID: data.IplBillID,
		ItemID:    data.ItemID,
		Note:      data.Note,
		SubAmount: data.SubAmount,
	}
}

func IplBillDetailCreateRequestToIplBillDetailModel(data IplBillDetailCreateRequest) models.IplBillDetail {
	return models.IplBillDetail{
		IplBillID: data.IplBillID,
		ItemID:    data.ItemID,
		Note:      data.Note,
		SubAmount: data.SubAmount,
	}
}
