package request

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type PostCommentCreateRequest struct {
	PostID  string `json:"post_id"`
	UserID  string `json:"user_id"`
	Comment string `json:"comment"`
}

type PostCommentUpdateRequset struct {
	ID string `json:"id"`
	PostCommentCreateRequest
}

func PostCommentUpdateRequsetToPostCommentModel(data PostCommentUpdateRequset) models.PostComment {
	base := models.BaseModel{
		ID: data.ID,
	}
	return models.PostComment{
		BaseModel: base,
		PostID:    data.PostID,
		UserID:    data.UserID,
		Comment:   data.Comment,
	}
}

func PostCommentCreateRequestToPostCommentModel(data PostCommentCreateRequest) models.PostComment {
	return models.PostComment{
		PostID:  data.PostID,
		UserID:  data.UserID,
		Comment: data.Comment,
	}
}
