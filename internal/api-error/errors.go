package api_error

import "fmt"

type Error struct {
	Code int
	Msg  string
}

func NewError(code int, msg string) *Error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("(%d) %s", e.Code, e.Msg)
}
