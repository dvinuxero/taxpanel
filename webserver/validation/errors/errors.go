package errors

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiError struct {
	message string
	error   string
	cause   string
	Status  int
}

// apiError implements the error interface
func (e ApiError) Error() string {
	return e.message
}

func (e ApiError) toMap() map[string]interface{} {
	return map[string]interface{}{
		"message": e.message,
		"error":   e.error,
		"status":  e.Status,
		"cause":   e.cause,
	}
}

// exported function that takes a a gin context and an error and makes the corresponding response
func RespondError(c *gin.Context, e error) {
	c.JSON(e.(ApiError).Status, e.(ApiError).toMap())
}

// aid functions to create different errors of type apiError
// they all return apiError values and only differ in attributes
func ValidationApiError(error, cause string) error {
	return ApiError{"Validation errors", error, cause, http.StatusBadRequest}
}

func NotFoundApiError(message string) error {
	return ApiError{message, "not_found", "", http.StatusNotFound}
}

func BadRequestApiError(message string) error {
	return ApiError{message, "bad_request", "", http.StatusBadRequest}
}

func InternalServerApiError(message string, err error) error {
	return ApiError{message, "internal_server_error", "", http.StatusInternalServerError}
}
