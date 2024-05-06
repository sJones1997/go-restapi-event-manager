package HTTPError

type HTTPError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func (e *HTTPError) Error() string {
	return e.Message
}

func New(statusCode int, message string) *HTTPError {
	return &HTTPError{StatusCode: statusCode, Message: message}
}
