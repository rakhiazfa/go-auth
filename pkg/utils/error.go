package utils

type HttpError struct {
	StatusCode int
	Message    string
	Reason     error
}

func (e *HttpError) Error() string {
	return e.Reason.Error()
}

func NewHttpError(statusCode int, message string, reason error) *HttpError {
	return &HttpError{statusCode, message, reason}
}

func CatchError(err error) {
	if err != nil {
		panic(err)
	}
}
