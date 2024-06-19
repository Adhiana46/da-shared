package errors

import "net/http"

type ErrorDuplicateEmail struct {
	msg string
}

func NewErrorDuplicateEmail(msg ...string) error {
	e := ErrorDuplicateEmail{}

	if len(msg) > 0 {
		e.msg = msg[0]
	}

	return &e
}

func (e *ErrorDuplicateEmail) Error() string {
	if e.msg != "" {
		return e.msg
	}

	return "email already exists"
}

func (e *ErrorDuplicateEmail) HttpStatusCode() int {
	return http.StatusBadRequest
}
