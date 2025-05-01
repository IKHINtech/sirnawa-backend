package request

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type VerifyEmailCodeRequest struct {
	Email string `json:"email" validate:"required,email"`
	Code  string `json:"code" validate:"required"`
}

type VerifyCodeRequest struct {
	Email string `json:"email" validate:"required,email"`
}
