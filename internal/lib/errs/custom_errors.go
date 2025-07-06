package errs

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

var (
	ErrTaskNotFound = errors.New("task not found")
	ErrAccessDenied = errors.New("access denied")
)

var (
	ErrEmptyCategory = errors.New("category empty")
)

var (
	ErrInvalidEmailOrPassword = errors.New("invalid email or password")
	ErrEmailIsAlreadyExists   = errors.New("email already exists")
)

func IsDuplicateError(err error) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == "23505"
}
