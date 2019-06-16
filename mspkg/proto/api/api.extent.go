package api

func (em *ErrorMessage) Code() string {
	return em.ErrCode
}

func (em *ErrorMessage) Error() string {
	return em.ErrMessage
}

// Returns the error details message.
func (em *ErrorMessage) Message() string {
	return em.ErrMessage
}

func (em *ErrorMessage) OrigErr() error {
	return nil
}
