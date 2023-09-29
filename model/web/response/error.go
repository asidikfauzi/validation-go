package response

type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       string `json:"data"`
}

type ErrorValidate struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
