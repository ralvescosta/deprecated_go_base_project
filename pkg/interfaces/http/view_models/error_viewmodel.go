package viewmodels

type ErrorMessage struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func StringToErrorResponse(message string) ErrorMessage {
	return ErrorMessage{
		Message: message,
	}
}
