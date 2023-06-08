package util_entity

type Validator interface {
	ValidateStruct(input interface{}) error
	GetValidationErrors() []string
}
