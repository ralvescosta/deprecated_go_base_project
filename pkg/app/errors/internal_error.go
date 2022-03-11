package errors

type InternalError struct {
	Message string
}

func (pst InternalError) Error() string {
	return pst.Message
}

func NewInternalError(m string) InternalError {
	return InternalError{Message: m}
}
