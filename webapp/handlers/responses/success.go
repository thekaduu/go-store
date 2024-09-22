package responses

type SuccessResponse struct {
	Message string         `json:"message"`
	Result  map[string]any `json:"data"`
}
