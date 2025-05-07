package request

import (
	"github.com/IKHINtech/sirnawa-backend/internal/models"
)

type AnnouncementCreateRequest struct {
	RtID        string   `json:"rt_id"`
	Title       string   `json:"title"`
	Content     string   `json:"content"`
	CreatedBy   string   `json:"created_by"`
	Attachments []string `json:"attachments"`
}

type AnnouncementUpdateRequset struct {
	ID string `json:"id"`
	AnnouncementCreateRequest
	DeleteAttachments []string `json:"delete_attachments"`
}

func AnnouncementUpdateRequsetToAnnouncementModel(data AnnouncementUpdateRequset) models.Announcement {
	base := models.BaseModel{
		ID: data.ID,
	}
	return models.Announcement{
		BaseModel:   base,
		RtID:        data.RtID,
		Title:       data.Title,
		Content:     data.Content,
		CreatedBy:   data.CreatedBy,
		Attachments: data.Attachments,
	}
}

func AnnouncementCreateRequestToAnnouncementModel(data AnnouncementCreateRequest) models.Announcement {
	return models.Announcement{
		RtID:        data.RtID,
		Title:       data.Title,
		Content:     data.Content,
		CreatedBy:   data.CreatedBy,
		Attachments: data.Attachments,
	}
}
