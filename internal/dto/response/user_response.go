package response

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type UserResponse struct {
	ID         string  `json:"id"`
	Email      *string `json:"email"`
	Role       string  `json:"role"`
	ResidentID *string `json:"resident_id"`
}

type UserDetailResponse struct {
	UserResponse
	UserRts  []UserRtResponse        `json:"user_rt"`
	Resident *ResidentDetailResponse `json:"resident"`
}

func UserToResponse(user *models.User) *UserDetailResponse {
	if user == nil {
		return nil
	}
	userRts := make([]UserRtResponse, len(user.UserRTs))

	for i, userRt := range user.UserRTs {
		userRts[i] = *UserRtToResponse(&userRt)
	}
	return &UserDetailResponse{
		UserResponse: UserResponse{
			ID:         user.ID,
			Email:      user.Email,
			Role:       string(user.Role),
			ResidentID: user.ResidentID,
		},
		UserRts:  userRts,
		Resident: MapResidentDetailResponse(user.Resident),
	}
}
