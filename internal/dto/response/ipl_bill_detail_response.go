package response

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type IplBillDetailResponse struct {
	BaseResponse
	IplBillID string `json:"ipl_bill_id"`
	ItemID    string `json:"item_id"`
	Note      string `json:"note"`
	SubAmount int64  `json:"sub_amount"`
}

type IplBillDetailResponses []IplBillDetailResponse

func IplBillDetailModelToIplBillDetailResponse(data *models.IplBillDetail) *IplBillDetailResponse {
	base := BaseResponse{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	return &IplBillDetailResponse{
		BaseResponse: base,
		IplBillID:    data.IplBillID,
		ItemID:       data.ItemID,
		Note:         data.Note,
		SubAmount:    data.SubAmount,
	}
}

func IplBillDetailListToResponse(data models.IplBillDetails) IplBillDetailResponses {
	var res IplBillDetailResponses
	for _, v := range data {
		res = append(res, *IplBillDetailModelToIplBillDetailResponse(&v))
	}
	return res
}
