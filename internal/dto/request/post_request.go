package request

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type PostCreateRequest struct {
	UserID      string   `json:"user_id"`
	Title       string   `json:"title"`
	Content     string   `json:"content"`
	Attachments []string `json:"attachments"`
}

type PostUpdateRequset struct {
	ID string `json:"id"`
	PostCreateRequest
}

func PostUpdateRequsetToPostModel(data PostUpdateRequset) models.Post {
	base := models.BaseModel{
		ID: data.ID,
	}
	return models.Post{
		BaseModel: base,
	}
}

func PostCreateRequestToPostModel(data PostCreateRequest) models.Post {
	return models.Post{
		UserID:      data.UserID,
		Title:       data.Title,
		Content:     data.Content,
		Attachments: data.Attachments,
	}
}
