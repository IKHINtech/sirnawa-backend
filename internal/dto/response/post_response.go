package response

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type PostResponse struct {
	BaseResponse
	UserID      string   `json:"user_id"`
	Title       string   `json:"title"`
	Content     string   `json:"content"`
	Attachments []string `json:"attachments"`
}

type PostResponses []PostResponse

func PostModelToPostResponse(data *models.Post) *PostResponse {
	base := BaseResponse{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	return &PostResponse{
		UserID:       data.UserID,
		Title:        data.Title,
		Content:      data.Content,
		Attachments:  data.Attachments,
		BaseResponse: base,
	}
}

func PostListToResponse(data models.Posts) PostResponses {
	var res PostResponses
	for _, v := range data {
		res = append(res, *PostModelToPostResponse(&v))
	}
	return res
}
