package response

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type BlockResponse struct {
	BaseResponse
	Name string `json:"name"`
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
