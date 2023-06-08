package utils

import (
	"net/http"

	"github.com/IbnAnjung/datting/entity/util_entity"
)

type AppError interface {
	Error() string
	ErrorCode() int
}
type ClientError struct {
	Message string
	Code    int
}

func (e ClientError) Error() string {
	return e.Message
}

func (e ClientError) ErrorCode() int {
	return e.Code
}

var (
	DataNotFoundError   = ClientError{Message: "Data Not Found", Code: http.StatusNotFound}
	DuplicatedDataError = ClientError{Message: "Data Already Exists", Code: http.StatusBadRequest}
)

type ValidationError struct {
	Message   string
	Validator util_entity.Validator
}

func (e ValidationError) Error() string {
	return e.Message
}

func (e ValidationError) ErrorCode() int {
	return http.StatusBadRequest
}

type ServerError struct {
}

func (e ServerError) Error() string {
	return "interval server error"
}

func (e ServerError) ErrorCode() int {
	return http.StatusInternalServerError
}
