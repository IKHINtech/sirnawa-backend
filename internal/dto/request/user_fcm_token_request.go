package request

type RemoveTokenRequest struct {
	Token string `json:"token" validate:"required"`
}

type RegisterTokenRequest struct {
	Token      string `json:"token" validate:"required"`
	DeviceID   string `json:"device_id" validate:"required"`
	DeviceType string `json:"device_type" validate:"required,oneof=android ios web other"`
	AppVersion string `json:"app_version"`
	OSVersion  string `json:"os_version"`
}
