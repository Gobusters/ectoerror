package httperror

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHTTPError(t *testing.T) {
	t.Run("creates new HTTPError", func(t *testing.T) {
		err := NewHTTPError(http.StatusBadRequest, "bad request")
		assert.Equal(t, http.StatusBadRequest, err.Code)
		assert.Equal(t, "bad request", err.Message)
		assert.NotNil(t, err.Meta)
	})
}

func TestNewHTTPErrorf(t *testing.T) {
	t.Run("creates new HTTPError with formatted message", func(t *testing.T) {
		err := NewHTTPErrorf(http.StatusNotFound, "user %d not found", 123)
		assert.Equal(t, http.StatusNotFound, err.Code)
		assert.Equal(t, "user 123 not found", err.Message)
	})
}

func TestWrapError(t *testing.T) {
	t.Run("wraps standard error", func(t *testing.T) {
		stdErr := errors.New("standard error")
		httpErr := WrapError(http.StatusBadRequest, stdErr)
		assert.Equal(t, http.StatusBadRequest, httpErr.Code)
		assert.Equal(t, "standard error", httpErr.Message)
		assert.Equal(t, stdErr, httpErr.err)
	})

	t.Run("returns existing HTTPError", func(t *testing.T) {
		existingErr := NewHTTPError(http.StatusNotFound, "not found")
		httpErr := WrapError(http.StatusBadRequest, existingErr)
		assert.Equal(t, existingErr, httpErr)
	})

	t.Run("uses default status code", func(t *testing.T) {
		stdErr := errors.New("standard error")
		httpErr := WrapError(0, stdErr)
		assert.Equal(t, http.StatusInternalServerError, httpErr.Code)
	})
}

func TestHTTPErrorError(t *testing.T) {
	t.Run("returns formatted error string", func(t *testing.T) {
		err := NewHTTPError(http.StatusBadRequest, "bad request")
		expected := "[400] HTTP Error: - bad request"
		assert.Equal(t, expected, err.Error())
	})
}

func TestHTTPErrorAddMetaValue(t *testing.T) {
	t.Run("adds meta value", func(t *testing.T) {
		err := NewHTTPError(http.StatusBadRequest, "bad request")
		err.AddMetaValue("key", "value")
		assert.Equal(t, "value", err.Meta["key"])
	})
}

func TestGetStatusCode(t *testing.T) {
	t.Run("returns status code for HTTPError", func(t *testing.T) {
		err := NewHTTPError(http.StatusNotFound, "not found")
		assert.Equal(t, http.StatusNotFound, GetStatusCode(err))
	})

	t.Run("returns 500 for non-HTTPError", func(t *testing.T) {
		err := errors.New("standard error")
		assert.Equal(t, http.StatusInternalServerError, GetStatusCode(err))
	})
}

func TestIsHTTPError(t *testing.T) {
	t.Run("returns true for HTTPError", func(t *testing.T) {
		err := NewHTTPError(http.StatusBadRequest, "bad request")
		assert.True(t, IsHTTPError(err))
	})

	t.Run("returns false for non-HTTPError", func(t *testing.T) {
		err := errors.New("standard error")
		assert.False(t, IsHTTPError(err))
	})
}

func TestIsStatus(t *testing.T) {
	t.Run("returns true for matching status", func(t *testing.T) {
		err := NewHTTPError(http.StatusTeapot, "I'm a teapot")
		assert.True(t, IsStatus(err, http.StatusTeapot))
	})

	t.Run("returns false for non-matching status", func(t *testing.T) {
		err := NewHTTPError(http.StatusOK, "OK")
		assert.False(t, IsStatus(err, http.StatusBadRequest))
	})
}

// Add test functions for IsOK, IsCreated, IsAccepted, etc.
func TestIsOK(t *testing.T) {
	t.Run("returns true for OK status", func(t *testing.T) {
		err := NewHTTPError(http.StatusOK, "OK")
		assert.True(t, IsOK(err))
	})

	t.Run("returns false for non-OK status", func(t *testing.T) {
		err := NewHTTPError(http.StatusBadRequest, "Bad Request")
		assert.False(t, IsOK(err))
	})
}

func TestIsCreated(t *testing.T) {
	t.Run("returns true for Created status", func(t *testing.T) {
		err := NewHTTPError(http.StatusCreated, "Created")
		assert.True(t, IsCreated(err))
	})

	t.Run("returns false for non-Created status", func(t *testing.T) {
		err := NewHTTPError(http.StatusOK, "OK")
		assert.False(t, IsCreated(err))
	})
}

func TestIsAccepted(t *testing.T) {
	t.Run("returns true for Accepted status", func(t *testing.T) {
		err := NewHTTPError(http.StatusAccepted, "Accepted")
		assert.True(t, IsAccepted(err))
	})

	t.Run("returns false for non-Accepted status", func(t *testing.T) {
		err := NewHTTPError(http.StatusOK, "OK")
		assert.False(t, IsAccepted(err))
	})
}

func TestIsNoContent(t *testing.T) {
	t.Run("returns true for NoContent status", func(t *testing.T) {
		err := NewHTTPError(http.StatusNoContent, "No Content")
		assert.True(t, IsNoContent(err))
	})

	t.Run("returns false for non-NoContent status", func(t *testing.T) {
		err := NewHTTPError(http.StatusOK, "OK")
		assert.False(t, IsNoContent(err))
	})
}

func TestIsBadRequest(t *testing.T) {
	t.Run("returns true for BadRequest status", func(t *testing.T) {
		err := NewHTTPError(http.StatusBadRequest, "Bad Request")
		assert.True(t, IsBadRequest(err))
	})

	t.Run("returns false for non-BadRequest status", func(t *testing.T) {
		err := NewHTTPError(http.StatusOK, "OK")
		assert.False(t, IsBadRequest(err))
	})
}

func TestIsUnauthorized(t *testing.T) {
	t.Run("returns true for Unauthorized status", func(t *testing.T) {
		err := NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		assert.True(t, IsUnauthorized(err))
	})

	t.Run("returns false for non-Unauthorized status", func(t *testing.T) {
		err := NewHTTPError(http.StatusOK, "OK")
		assert.False(t, IsUnauthorized(err))
	})
}

func TestIsForbidden(t *testing.T) {
	t.Run("returns true for Forbidden status", func(t *testing.T) {
		err := NewHTTPError(http.StatusForbidden, "Forbidden")
		assert.True(t, IsForbidden(err))
	})

	t.Run("returns false for non-Forbidden status", func(t *testing.T) {
		err := NewHTTPError(http.StatusOK, "OK")
		assert.False(t, IsForbidden(err))
	})
}

func TestIsNotFound(t *testing.T) {
	t.Run("returns true for NotFound status", func(t *testing.T) {
		err := NewHTTPError(http.StatusNotFound, "Not Found")
		assert.True(t, IsNotFound(err))
	})

	t.Run("returns false for non-NotFound status", func(t *testing.T) {
		err := NewHTTPError(http.StatusOK, "OK")
		assert.False(t, IsNotFound(err))
	})
}

func TestIsInternalServerError(t *testing.T) {
	t.Run("returns true for InternalServerError status", func(t *testing.T) {
		err := NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
		assert.True(t, IsInternalServerError(err))
	})

	t.Run("returns false for non-InternalServerError status", func(t *testing.T) {
		err := NewHTTPError(http.StatusOK, "OK")
		assert.False(t, IsInternalServerError(err))
	})
}

func TestIsBadGateway(t *testing.T) {
	t.Run("returns true for BadGateway status", func(t *testing.T) {
		err := NewHTTPError(http.StatusBadGateway, "Bad Gateway")
		assert.True(t, IsBadGateway(err))
	})

	t.Run("returns false for non-BadGateway status", func(t *testing.T) {
		err := NewHTTPError(http.StatusOK, "OK")
		assert.False(t, IsBadGateway(err))
	})
}

func TestIsServiceUnavailable(t *testing.T) {
	t.Run("returns true for ServiceUnavailable status", func(t *testing.T) {
		err := NewHTTPError(http.StatusServiceUnavailable, "Service Unavailable")
		assert.True(t, IsServiceUnavailable(err))
	})

	t.Run("returns false for non-ServiceUnavailable status", func(t *testing.T) {
		err := NewHTTPError(http.StatusOK, "OK")
		assert.False(t, IsServiceUnavailable(err))
	})
}

func TestIsGatewayTimeout(t *testing.T) {
	t.Run("returns true for GatewayTimeout status", func(t *testing.T) {
		err := NewHTTPError(http.StatusGatewayTimeout, "Gateway Timeout")
		assert.True(t, IsGatewayTimeout(err))
	})

	t.Run("returns false for non-GatewayTimeout status", func(t *testing.T) {
		err := NewHTTPError(http.StatusOK, "OK")
		assert.False(t, IsGatewayTimeout(err))
	})
}

func TestIsError(t *testing.T) {
	t.Run("returns true for error status codes", func(t *testing.T) {
		errorCodes := []int{http.StatusBadRequest, http.StatusInternalServerError, http.StatusNotImplemented}
		for _, code := range errorCodes {
			err := NewHTTPError(code, "Error")
			assert.True(t, IsError(err))
		}
	})

	t.Run("returns false for non-error status codes", func(t *testing.T) {
		nonErrorCodes := []int{http.StatusOK, http.StatusCreated, http.StatusNoContent}
		for _, code := range nonErrorCodes {
			err := NewHTTPError(code, "Non-Error")
			assert.False(t, IsError(err))
		}
	})
}

func TestIsSuccess(t *testing.T) {
	t.Run("returns true for success status codes", func(t *testing.T) {
		successCodes := []int{http.StatusOK, http.StatusCreated, http.StatusAccepted}
		for _, code := range successCodes {
			err := NewHTTPError(code, "Success")
			assert.True(t, IsSuccess(err))
		}
	})

	t.Run("returns false for non-success status codes", func(t *testing.T) {
		nonSuccessCodes := []int{http.StatusBadRequest, http.StatusInternalServerError, http.StatusNotFound}
		for _, code := range nonSuccessCodes {
			err := NewHTTPError(code, "Non-Success")
			assert.False(t, IsSuccess(err))
		}
	})
}

func TestIsRedirect(t *testing.T) {
	t.Run("returns true for redirect status codes", func(t *testing.T) {
		redirectCodes := []int{http.StatusMultipleChoices, http.StatusMovedPermanently, http.StatusTemporaryRedirect}
		for _, code := range redirectCodes {
			err := NewHTTPError(code, "Redirect")
			assert.True(t, IsRedirect(err))
		}
	})

	t.Run("returns false for non-redirect status codes", func(t *testing.T) {
		nonRedirectCodes := []int{http.StatusOK, http.StatusBadRequest, http.StatusInternalServerError}
		for _, code := range nonRedirectCodes {
			err := NewHTTPError(code, "Non-Redirect")
			assert.False(t, IsRedirect(err))
		}
	})
}

func TestHTTPErrorUnwrap(t *testing.T) {
	t.Run("unwraps to original error", func(t *testing.T) {
		originalErr := errors.New("original error")
		httpErr := WrapError(http.StatusBadRequest, originalErr)
		assert.Equal(t, originalErr, errors.Unwrap(httpErr))
	})

	t.Run("unwraps to nil for non-wrapped error", func(t *testing.T) {
		httpErr := NewHTTPError(http.StatusBadRequest, "bad request")
		assert.Nil(t, errors.Unwrap(httpErr))
	})
}

func TestIsClientError(t *testing.T) {
	t.Run("returns true for client error status codes", func(t *testing.T) {
		clientErrorCodes := []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden}
		for _, code := range clientErrorCodes {
			err := NewHTTPError(code, "Client Error")
			assert.True(t, IsClientError(err))
		}
	})

	t.Run("returns false for non-client error status codes", func(t *testing.T) {
		nonClientErrorCodes := []int{http.StatusOK, http.StatusInternalServerError, http.StatusPermanentRedirect}
		for _, code := range nonClientErrorCodes {
			err := NewHTTPError(code, "Non-Client Error")
			assert.False(t, IsClientError(err))
		}
	})
}

func TestIsServerError(t *testing.T) {
	t.Run("returns true for server error status codes", func(t *testing.T) {
		serverErrorCodes := []int{http.StatusInternalServerError, http.StatusBadGateway, http.StatusServiceUnavailable}
		for _, code := range serverErrorCodes {
			err := NewHTTPError(code, "Server Error")
			assert.True(t, IsServerError(err))
		}
	})

	t.Run("returns false for non-server error status codes", func(t *testing.T) {
		nonServerErrorCodes := []int{http.StatusOK, http.StatusBadRequest}
		for _, code := range nonServerErrorCodes {
			err := NewHTTPError(code, "Non-Server Error")
			assert.False(t, IsServerError(err))
		}
	})
}
