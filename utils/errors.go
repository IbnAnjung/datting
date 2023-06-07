package utils

import (
	"github.com/IbnAnjung/datting/entity/validator_entity"
)

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
	DataNotFoundError   = ClientError{Message: "Data Not Found"}
	DuplicatedDataError = ClientError{Message: "Data Already Exists"}
)

type ValidationError struct {
	Message   string
	Validator validator_entity.Validator
}

func (e ValidationError) Error() string {
	return e.Message
}
