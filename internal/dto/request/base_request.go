package request

type BaseCreateRequest struct{}

type BaseUpdateRequset struct {
	ID string `json:"id"`
	BaseCreateRequest
}
