package http

type ErrorType uint64

type Error struct {
	Err  error
	Type ErrorType
	Meta interface{}
}

type errorMsgs []*Error

var _ error = &Error{}

func (msg *Error) Error() string {
	return msg.Err.Error()
}
