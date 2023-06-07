package utils

type DataNotFoundError struct{}

func (DataNotFoundError) Error() string {
	return "Data Not Found"
}

type DuplicatedDataError struct{}

func (DuplicatedDataError) Error() string {
	return "Data already exists"
}

type ValidationError struct {
	Message string
}

func (e ValidationError) Error() string {
	return e.Message
}
