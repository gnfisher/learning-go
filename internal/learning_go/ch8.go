package learning_go

import (
	"errors"
	"fmt"
)

func MakingErrors() error {
	return errors.New("my error")
}

type Thing interface{}

func InterfaceIsNil() Thing {
	var t Thing
	return t
}

func PlayWithErrors() {
	err := MakingErrors()
	if err != nil {
		fmt.Println(err)
		fmt.Println(fmt.Errorf("a new error message: %w", err))
		fmt.Println(err)
	}
}

func NewError(text string) error {
	return &errorString{text}
}

// errorString is a trivial implementation of error.
// the error interface requires a method, `Error()` that returns string.
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

// Errors with context?

type errorWithContext struct {
	s string
	i int
}

func (e *errorWithContext) Error() string {
	return e.s
}

func (e *errorWithContext) StatusCode() int {
	return e.i
}

// Sentinel Error
// A type that implements the error interface
// Lets us do == checks (but requires strings match)

type SentinelError string

func (s SentinelError) Error() string {
	return string(s)
}
