package applicationerrors

import "net/http"

type ApplicationError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     string `json:"error"`
}

func NewApplicationError(code int, message, err string) *ApplicationError {

	return &ApplicationError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func NotFoundException(message string) *ApplicationError {
	return NewApplicationError(http.StatusNotFound, message, "NOT_FOUND_EXCEPTION")
}

func InternalServerException() *ApplicationError {
	return NewApplicationError(http.StatusNotFound, "", "INTERNAL_SERVER_EXCEPTION")
}

func ForbiddenException() *ApplicationError {
	return NewApplicationError(http.StatusForbidden, "", "FORBIDDEN_EXCEPTION")
}

func BadRequestException(message string) *ApplicationError {
	return NewApplicationError(http.StatusBadRequest, message, "BAD_REQUEST_EXCEPTION")
}
