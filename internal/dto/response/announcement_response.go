package response

import (
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/lib/pq"
)

type AnnouncementResponse struct {
	BaseResponse
	RtID        string         `json:"rt_id"`
	Title       string         `json:"title"`
	Content     string         `json:"content"`
	CreatedBy   string         `json:"created_by"`
	Attachments pq.StringArray `json:"attachments"`
}

type AnnouncementResponses []AnnouncementResponse

func AnnouncementModelToAnnouncementResponse(data *models.Announcement) *AnnouncementResponse {
	base := BaseResponse{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	return &AnnouncementResponse{
		BaseResponse: base,
		RtID:         data.RtID,
		Title:        data.Title,
		Content:      data.Content,
		CreatedBy:    data.CreatedBy,
		Attachments:  data.Attachments,
	}
}

func AnnouncementListToResponse(data models.Announcements) AnnouncementResponses {
	var res AnnouncementResponses
	for _, v := range data {
		res = append(res, *AnnouncementModelToAnnouncementResponse(&v))
	}
	return res
}
