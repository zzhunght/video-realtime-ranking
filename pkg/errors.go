package errors

import "errors"

var (
	ErrNotfound       = errors.New("resource not found")
	ErrInternalServer = errors.New("internal server error")
	ErrBadrequest     = errors.New("bad request")
)
