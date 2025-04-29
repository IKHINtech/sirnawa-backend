package response

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type UserRtResponse struct {
	BaseResponse
	UserID string     `json:"user_id"`
	RtID   string     `json:"rt_id"`
	Rt     RtResponse `json:"rt"`
	Role   string     `json:"role"`
}

func UserRtToResponse(data *models.UserRT) *UserRtResponse {
	if data == nil {
		return nil
	}
	return &UserRtResponse{
		BaseResponse: BaseResponse{
			ID: data.ID,
		},
		UserID: data.ID,
		RtID:   data.RtID,
		Rt:     *RtModelToRtResponse(&data.Rt),
		Role:   string(data.Role),
	}
}
