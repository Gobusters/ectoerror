package httperror

import (
	"errors"
	"fmt"
	"net/http"
)

// HTTPError represents an error that occurred during an HTTP request.
// It contains the HTTP status code, a message, and optional metadata.
type HTTPError struct {
	Code    int
	Message string
	Meta    map[string]any
	err     error
}

// NewHTTPError creates a new HTTPError with the given status code and message.
func NewHTTPError(code int, message string) *HTTPError {
	return &HTTPError{Code: code, Message: message, Meta: make(map[string]any)}
}

// NewHTTPErrorf creates a new HTTPError with the given status code and formatted message.
func NewHTTPErrorf(code int, format string, args ...any) *HTTPError {
	message := fmt.Sprintf(format, args...)
	return &HTTPError{Code: code, Message: message, Meta: make(map[string]any)}
}

// Implement the Unwrap method
func (e *HTTPError) Unwrap() error {
	return e.err
}

// WrapError wraps an error with an HTTPError.
func WrapError(code int, err error) *HTTPError {
	if httpErr, ok := err.(*HTTPError); ok {
		return httpErr
	}
	if code == 0 {
		code = http.StatusInternalServerError
	}
	return &HTTPError{Code: code, Message: err.Error(), Meta: make(map[string]any), err: err}
}

// Error returns the error message as a string.
func (e *HTTPError) Error() string {
	return fmt.Sprintf("[%d] HTTP Error: - %s", e.Code, e.Message)
}

// AddMetaValue adds a metadata value to the HTTPError.
func (e *HTTPError) AddMetaValue(key string, value any) *HTTPError {
	e.Meta[key] = value
	return e
}

// GetStatusCode returns the HTTP status code from the provided error.
// If the error is not an HTTPError, it returns 500.
func GetStatusCode(err error) int {
	if e, ok := err.(*HTTPError); ok {
		return e.Code
	}
	return http.StatusInternalServerError
}

// Modify IsHTTPError to use errors.As
func IsHTTPError(err error) bool {
	var httpErr *HTTPError
	return errors.As(err, &httpErr)
}

// ToHTTPError converts an error to an HTTPError if it is an HTTPError.
func ToHTTPError(err error) *HTTPError {
	if err == nil {
		return nil
	}
	if httpErr, ok := err.(*HTTPError); ok {
		return httpErr
	}
	return WrapError(GetStatusCode(err), err)
}

// IsOK checks if the provided error is an HTTPError with a status code of 200.
func IsOK(err error) bool {
	var httpErr *HTTPError
	return errors.As(err, &httpErr) && httpErr.Code == http.StatusOK
}

// IsCreated checks if the provided error is an HTTPError with a status code of 201.
func IsCreated(err error) bool {
	var httpErr *HTTPError
	return errors.As(err, &httpErr) && httpErr.Code == http.StatusCreated
}

// IsAccepted checks if the provided error is an HTTPError with a status code of 202.
func IsAccepted(err error) bool {
	var httpErr *HTTPError
	return errors.As(err, &httpErr) && httpErr.Code == http.StatusAccepted
}

// IsNoContent checks if the provided error is an HTTPError with a status code of 204.
func IsNoContent(err error) bool {
	var httpErr *HTTPError
	return errors.As(err, &httpErr) && httpErr.Code == http.StatusNoContent
}

// IsBadRequest checks if the provided error is an HTTPError with a status code of 400.
func IsBadRequest(err error) bool {
	var httpErr *HTTPError
	return errors.As(err, &httpErr) && httpErr.Code == http.StatusBadRequest
}

// IsUnauthorized checks if the provided error is an HTTPError with a status code of 401.
func IsUnauthorized(err error) bool {
	var httpErr *HTTPError
	return errors.As(err, &httpErr) && httpErr.Code == http.StatusUnauthorized
}

// IsForbidden checks if the provided error is an HTTPError with a status code of 403.
func IsForbidden(err error) bool {
	var httpErr *HTTPError
	return errors.As(err, &httpErr) && httpErr.Code == http.StatusForbidden
}

// IsNotFound checks if the provided error is an HTTPError with a status code of 404.
func IsNotFound(err error) bool {
	var httpErr *HTTPError
	return errors.As(err, &httpErr) && httpErr.Code == http.StatusNotFound
}

// IsInternalServerError checks if the provided error is an HTTPError with a status code of 500.
func IsInternalServerError(err error) bool {
	var httpErr *HTTPError
	return errors.As(err, &httpErr) && httpErr.Code == http.StatusInternalServerError
}

// IsBadGateway checks if the provided error is an HTTPError with a status code of 502.
func IsBadGateway(err error) bool {
	var httpErr *HTTPError
	return errors.As(err, &httpErr) && httpErr.Code == http.StatusBadGateway
}

// IsServiceUnavailable checks if the provided error is an HTTPError with a status code of 503.
func IsServiceUnavailable(err error) bool {
	var httpErr *HTTPError
	return errors.As(err, &httpErr) && httpErr.Code == http.StatusServiceUnavailable
}

// IsGatewayTimeout checks if the provided error is an HTTPError with a status code of 504.
func IsGatewayTimeout(err error) bool {
	var httpErr *HTTPError
	return errors.As(err, &httpErr) && httpErr.Code == http.StatusGatewayTimeout
}

// IsStatus checks if the provided error is an HTTPError with the specified status code.
func IsStatus(err error, status int) bool {
	var httpErr *HTTPError
	return errors.As(err, &httpErr) && httpErr.Code == status
}

// IsClientError checks if the provided error is an HTTPError with a status code of 4XX.
func IsClientError(err error) bool {
	var httpErr *HTTPError
	return errors.As(err, &httpErr) && httpErr.Code >= http.StatusBadRequest && httpErr.Code < http.StatusInternalServerError
}

// IsServerError checks if the provided error is an HTTPError with a status code of 5XX.
func IsServerError(err error) bool {
	var httpErr *HTTPError
	return errors.As(err, &httpErr) && httpErr.Code >= http.StatusInternalServerError && httpErr.Code < http.StatusNetworkAuthenticationRequired
}

// IsError checks if the provided error is an HTTPError with a status code of 400 or higher.
func IsError(err error) bool {
	var httpErr *HTTPError
	return errors.As(err, &httpErr) && httpErr.Code >= http.StatusBadRequest && httpErr.Code < http.StatusNetworkAuthenticationRequired
}

// IsSuccess checks if the provided error is an HTTPError with a status code of 200.
func IsSuccess(err error) bool {
	var httpErr *HTTPError
	return errors.As(err, &httpErr) && httpErr.Code >= http.StatusOK && httpErr.Code < http.StatusMultipleChoices
}

// IsRedirect checks if the provided error is an HTTPError with a status code of 3XX.
func IsRedirect(err error) bool {
	var httpErr *HTTPError
	return errors.As(err, &httpErr) && httpErr.Code >= http.StatusMultipleChoices && httpErr.Code < http.StatusBadRequest
}
