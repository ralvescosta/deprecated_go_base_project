package errors

type NotFoundError struct {
	Message string
}

func (pst NotFoundError) Error() string {
	return pst.Message
}

func NewNotFoundError(message string) NotFoundError {
	return NotFoundError{Message: message}
}
