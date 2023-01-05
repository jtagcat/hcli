package harg

import (
	"errors"
)

type genericErr struct {
	Err     error
	Wrapped error
}

func (a genericErr) Is(target error) bool {
	return errors.Is(a.Err, target)
}

func (a genericErr) Unwrap() error {
	return a.Wrapped
}

func (a genericErr) Error() string {
	return a.Err.Error() + ": " + a.Wrapped.Error()
}
