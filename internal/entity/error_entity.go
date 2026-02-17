package entity

import "errors"

var (
	ErrGlobalNotFound  = errors.New("not_found")
	ErrGlobalServerErr = errors.New("internal_server_error")
)
