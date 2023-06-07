package utils

import "github.com/IbnAnjung/datting/entity/validator_entity"

type DataNotFoundError struct{}

func (DataNotFoundError) Error() string {
	return "Data Not Found"
}

type DuplicatedDataError struct{}

func (DuplicatedDataError) Error() string {
	return "Data already exists"
}

type ValidationError struct {
	Message   string
	Validator validator_entity.Validator
}

func (e ValidationError) Error() string {
	return e.Message
}
