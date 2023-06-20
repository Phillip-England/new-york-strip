package core

import "fmt"

type HttpErr struct {
	Code int
	Message string
	Trace int
}

func NewHttpErr(trace int, code int, message string) (*HttpErr) {
	return &HttpErr{
		Code: code,
		Message: message,
		Trace: trace,
	}
}

func (httpErr *HttpErr) Error() string {
	return fmt.Sprintf("Trace: %d, Code: %d, Message: %s", httpErr.Trace, httpErr.Code, httpErr.Message)
}