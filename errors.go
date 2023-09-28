package server

type RequestError struct {
	s    string
	code int
}

func (err *RequestError) Error() string {
	return err.s
}

func (err *RequestError) ErrorCode() int {
	return err.code
}
