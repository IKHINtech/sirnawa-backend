package response

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type PostCommentResponse struct {
	BaseResponse
	PostID  string `json:"post_id"`
	UserID  string `json:"user_id"`
	Comment string `json:"comment"`
}

type PostCommentResponses []PostCommentResponse

func PostCommentModelToPostCommentResponse(data *models.PostComment) *PostCommentResponse {
	base := BaseResponse{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	return &PostCommentResponse{
		BaseResponse: base,
		PostID:       data.PostID,
		UserID:       data.UserID,
		Comment:      data.Comment,
	}
}

func PostCommentListToResponse(data models.PostComments) PostCommentResponses {
	var res PostCommentResponses
	for _, v := range data {
		res = append(res, *PostCommentModelToPostCommentResponse(&v))
	}
	return res
}
