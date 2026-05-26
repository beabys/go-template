package example

import "errors"

var (
	ErrHelloWorldNotFound = errors.New("hello world not found")
	ErrInvalidMessage     = errors.New("hello world message cannot be empty")
)
