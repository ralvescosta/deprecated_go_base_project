package errors

type ConflictError struct {
	Message string
}

func (pst ConflictError) Error() string {
	return pst.Message
}

func NewConflictError(message string) ConflictError {
	return ConflictError{Message: message}
}
