package core

type ErrServerFailure struct {
	Message string
	Code int
	Trace string
}

func NewErrServerFailure(message string, trace string) (*ErrServerFailure) {
	return &ErrServerFailure{
		Message: message,
		Code: 500,
		Trace: trace,
	}
}

func (e *ErrServerFailure) Error() (string) {
	return e.Message
}