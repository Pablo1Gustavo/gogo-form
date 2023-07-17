package domain

type RequestError struct {
	Code    int
	Err     error
	Details any
}

func (e *RequestError) Error() string {
	return e.Err.Error()
}
