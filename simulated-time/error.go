package simulated_time

import "fmt"

type ErrorMultiple struct {
	error
	multiple int
}

func (e *ErrorMultiple) Error() string {
	return fmt.Sprintf("multiple %d is invalid", e.multiple)
}
