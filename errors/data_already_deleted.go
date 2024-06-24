package errors

import "net/http"

type ErrorDataAlreadyDeleted struct {
	msg string
}

func NewErrorDataAlreadyDeleted(msg ...string) error {
	e := ErrorDataAlreadyDeleted{}

	if len(msg) > 0 {
		e.msg = msg[0]
	}

	return &e
}

func (e *ErrorDataAlreadyDeleted) Error() string {
	if e.msg != "" {
		return e.msg
	}

	return "data already deleted"
}

func (e *ErrorDataAlreadyDeleted) HttpStatusCode() int {
	return http.StatusBadRequest
}
