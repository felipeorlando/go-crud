package utils

import "errors"

var (
	// ErrBadRequest is a request-reading error
	ErrBadRequest = errors.New("Bad request")
	// ErrInternalServer is an internal problem
	ErrInternalServer = errors.New("Internal server error")
	// ErrNotFound is used when we can't find
	ErrNotFound = errors.New("Not found")
	// ErrMethodNotAllowed is used when a method is not allowed for an endpoint
	ErrMethodNotAllowed = errors.New("Method not allowed")
)
