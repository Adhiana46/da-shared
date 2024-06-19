package errors

import "net/http"

type ErrorRepositoryStore struct {
	// ORIGINAL ERROR
	err error
	msg string
}

func NewErrorRepositoryStore(err error, msg ...string) error {
	e := ErrorRepositoryStore{}

	if len(msg) > 0 {
		e.msg = msg[0]
	}

	return &e
}

func (e *ErrorRepositoryStore) Error() string {
	if e.msg != "" {
		return e.msg
	}

	return e.err.Error()
}

func (e *ErrorRepositoryStore) HttpStatusCode() int {
	return http.StatusInternalServerError
}
