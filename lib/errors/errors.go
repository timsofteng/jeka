package apperrors

import "errors"

var (
	ErrTooManyRequest = errors.New("too many requests")
	ErrNotExisted     = errors.New("not existed")
	ErrNotFound       = errors.New("not found")
	ErrExternal       = errors.New("external call error")
)
