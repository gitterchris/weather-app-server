package core

// ServerError represents the server error.
// Message is the user friendly message returned to the user
type ServerError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
	Error      error  `json:"-"`
}

// NewServerError instantiates a new ServerError object
func NewServerError(message string, statusCode int, err error) *ServerError {
	return &ServerError{
		Message:    message,
		StatusCode: statusCode,
		Error:      err,
	}
}
