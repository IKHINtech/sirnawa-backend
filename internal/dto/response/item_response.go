package response

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type ItemResponse struct {
	BaseResponse
	Name        string `json:"name"`
	RtID        string `json:"rt_id"`
	ItemType    string `json:"item_type"`
	Description string `json:"description"`
}

type ItemResponses []ItemResponse

func ItemModelToItemResponse(data *models.Item) *ItemResponse {
	if data == nil {
		return nil
	}
	base := BaseResponse{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	return &ItemResponse{
		Name:         data.Name,
		RtID:         data.RtID,
		BaseResponse: base,
		ItemType:     data.ItemType.ToString(),
		Description:  data.Description,
	}
}

func ItemListToResponse(data models.Items) ItemResponses {
	var res ItemResponses
	for _, v := range data {
		res = append(res, *ItemModelToItemResponse(&v))
	}
	return res
}
