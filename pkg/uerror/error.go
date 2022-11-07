package uerror

import (
	"fmt"
	"net/http"
)

type Error interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

type customError struct {
	ErrMessage string        `json:"message"`
	ErrStatus  int           `json:"status"`
	ErrError   string        `json:"uerror"`
	ErrCauses  []interface{} `json:"causes"`
}

func (e *customError) Message() string {
	return fmt.Sprintf("message: %s - status: %d - uerror: %s - causes: %v",
		e.ErrMessage, e.ErrStatus, e.ErrError, e.ErrCauses)
}

func (e *customError) Status() int {
	return e.ErrStatus
}

func (e *customError) Error() string {
	return e.ErrError
}

func (e *customError) Causes() []interface{} {
	return e.ErrCauses
}

func NewRestError(message string, status int, err string, causes []interface{}) Error {
	return &customError{
		ErrMessage: message,
		ErrStatus:  status,
		ErrError:   err,
		ErrCauses:  causes,
	}
}

func NewInternalServerError(message string, err error) Error {
	result := &customError{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "internal_server_error",
	}
	if err != nil {
		result.ErrCauses = append(result.ErrCauses, err.Error())
	}
	return result
}

func NewNotFoundError(message string) Error {
	return &customError{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		ErrError:   "not_found",
	}
}
