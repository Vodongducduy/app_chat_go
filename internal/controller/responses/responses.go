package responses

const (
	FAILED     = "Failed"
	SUCCESSFUL = "Successful"
)

type (
	JsonResponse  struct{}
	ErrorResponse struct {
		Message string `json:"message"`
		Code    string `json:"code"`
	}
)
