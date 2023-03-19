package api

type BadRequestError struct {
	Error interface{}
}

func NewBadRequestError(error interface{}) BadRequestError {
	return BadRequestError{
		Error: error,
	}
}
