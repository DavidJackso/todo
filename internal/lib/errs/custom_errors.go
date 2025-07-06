package errs

import "errors"

var (
	ErrInvaliEmailOrPassword = errors.New("invalid email or password")
	ErrEmailIsReady          = errors.New("email is ready")
	ErrTaskNotFound          = errors.New("task not found")
	ErrAccessDenied          = errors.New("access denied")
)
