package types

import "fmt"

type JsonResponse map[string]any

type AppError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	RawError   error  `json:"-"`
}

func (e *AppError) Error() string {
	if e.RawError == nil {
		return fmt.Sprintf("[%d] %s", e.StatusCode, e.Message)
	}
	return fmt.Sprintf("[%d] %s: %v", e.StatusCode, e.Message, e.RawError)
}

func NewAppError(statusCode int, message string, rawError error) *AppError {
	return &AppError{
		StatusCode: statusCode,
		Message:    message,
		RawError:   rawError,
	}
}
