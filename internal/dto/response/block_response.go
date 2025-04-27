package response

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type BlockResponse struct {
	BaseResponse
	Name string `json:"name"`
	RtID string `json:"rt_id"`
}

type BlockResponses []BlockResponse

func BlockModelToBlockResponse(data *models.Block) *BlockResponse {
	base := BaseResponse{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	return &BlockResponse{
		Name:         data.Name,
		RtID:         data.RtID,
		BaseResponse: base,
	}
}

func BlockListToResponse(data models.Blocks) BlockResponses {
	var res BlockResponses
	for _, v := range data {
		res = append(res, *BlockModelToBlockResponse(&v))
	}
	return res
}
