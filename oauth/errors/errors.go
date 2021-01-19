package errors

import (
	"errors"
	"net/http"
)

// RestErr the rest err interface for errors
type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

// NewError allows for specified/customized error reporting to the client
func NewError(msg string) error {
	return errors.New(msg)
}

// NewBadRequestError is for dynamic bad request errors
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

// NewNotFoundError is for dynamic bad request errors
func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "Interal server error",
	}
}
