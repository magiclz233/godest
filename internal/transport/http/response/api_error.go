package response

import (
	"errors"
	"net/http"
)

type APIError struct {
	HTTPStatus int
	Code       string
	Message    string
	Err        error
}

func (e *APIError) Error() string {
	if e == nil {
		return ""
	}
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.Message
}

func NewAPIError(status int, code, message string, err error) *APIError {
	return &APIError{
		HTTPStatus: status,
		Code:       code,
		Message:    message,
		Err:        err,
	}
}

func BadRequest(message string) *APIError {
	return NewAPIError(http.StatusBadRequest, CodeBadRequest, message, nil)
}

func Unauthorized(message string) *APIError {
	return NewAPIError(http.StatusUnauthorized, CodeUnauthorized, message, nil)
}

func Conflict(message string) *APIError {
	return NewAPIError(http.StatusConflict, CodeConflict, message, nil)
}

func Internal(message string, err error) *APIError {
	return NewAPIError(http.StatusInternalServerError, CodeInternal, message, err)
}

func NormalizeError(err error) *APIError {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr
	}
	return Internal("internal server error", err)
}
