package response

type UserResponse struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	Role       string `json:"role"`
	ResidentID string `json:"resident_id"`
}

type UserDetailResponse struct {
	UserResponse
	Resident ResidentResponse `json:"resident"`
}
