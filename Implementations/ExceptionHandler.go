package GoExceptionHandler

type GeneralException struct {
	ErrMessage string
}

// Error implements error
func (error GeneralException) Error() string {
	return error.ErrMessage
}

func (error GeneralException) FatalErrorWrapper() error {
	return GeneralException{error.ErrMessage}
}
