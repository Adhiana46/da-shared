package errors

import (
	"fmt"
	"net/http"
)

type ErrorFieldValue struct {
	field string
	value any
}

func NewErrorFieldValue(field string, value any) error {
	e := ErrorFieldValue{
		field: field,
		value: value,
	}

	return &e
}

func (e *ErrorFieldValue) Error() string {
	return fmt.Sprintf("invalid value for field '%s': %v", e.field, e.value)
}

func (e *ErrorFieldValue) HttpStatusCode() int {
	return http.StatusBadRequest
}
