package errors

import "net/http"

type ErrorSetPassword struct {
	// ORIGINAL ERROR
	err error
	msg string
}

func NewErrorSetPassword(err error, msg ...string) error {
	e := ErrorSetPassword{}

	if len(msg) > 0 {
		e.msg = msg[0]
	}

	return &e
}

func (e *ErrorSetPassword) Error() string {
	if e.msg != "" {
		return e.msg
	}

	return e.err.Error()
}

func (e *ErrorSetPassword) HttpStatusCode() int {
	return http.StatusInternalServerError
}
