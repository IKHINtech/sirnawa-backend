package dto

type ResponseLogin struct {
	User        any   `json:"user"`
	AccessToken Token `json:"access_token"`
}

type Token struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

type RegisterInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}
