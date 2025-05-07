package response

import (
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
)

type AnnouncementResponse struct {
	BaseResponse
	RtID           string   `json:"rt_id"`
	Title          string   `json:"title"`
	Content        string   `json:"content"`
	CreatedBy      string   `json:"created_by"`
	Attachments    []string `json:"attachments"`
	AttachmentUrls []string `json:"attachment_urls"`
}

type AnnouncementResponses []AnnouncementResponse

func AnnouncementModelToAnnouncementResponse(data *models.Announcement,
	driveService utils.DriveService,
) *AnnouncementResponse {
	base := BaseResponse{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	attachmentUrls := make([]string, len(data.Attachments))
	for i, attachment := range data.Attachments {
		attachmentUrls[i] = driveService.GetFileURL(attachment)
	}
	return &AnnouncementResponse{
		BaseResponse:   base,
		RtID:           data.RtID,
		Title:          data.Title,
		Content:        data.Content,
		CreatedBy:      data.CreatedBy,
		Attachments:    data.Attachments,
		AttachmentUrls: attachmentUrls,
	}
}

func AnnouncementListToResponse(data models.Announcements,
	driveService utils.DriveService,
) AnnouncementResponses {
	var res AnnouncementResponses
	for _, v := range data {
		res = append(res, *AnnouncementModelToAnnouncementResponse(&v, driveService))
	}
	return res
}
