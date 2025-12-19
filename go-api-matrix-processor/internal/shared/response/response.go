package response

// APIResponse defines the standard structure for JSON responses
// APIResponse defines the standard structure for JSON responses
type APIResponse[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
	Errors  any    `json:"errors,omitempty"`
}

// SuccessResponse creates a successful response
func SuccessResponse[T any](message string, data T) APIResponse[T] {
	return APIResponse[T]{
		Success: true,
		Message: message,
		Data:    data,
	}
}

// ErrorResponse creates an error response
func ErrorResponse(message string, errorDetail string) APIResponse[any] {
	var errDetails []string
	if errorDetail != "" {
		errDetails = append(errDetails, errorDetail)
	} else {
		// If no detail, use message as the error details per user request example
		errDetails = append(errDetails, message)
	}

	return APIResponse[any]{
		Success: false,
		Message: message,
		Errors:  errDetails,
	}
}
