package errors

import "net/http"

type ErrorUserNotFound struct {
	msg string
}

func NewErrorUserNotFound(msg ...string) error {
	e := ErrorUserNotFound{}

	if len(msg) > 0 {
		e.msg = msg[0]
	}

	return &e
}

func (e *ErrorUserNotFound) Error() string {
	if e.msg != "" {
		return e.msg
	}

	return "user not found"
}

func (e *ErrorUserNotFound) HttpStatusCode() int {
	return http.StatusBadRequest
}
