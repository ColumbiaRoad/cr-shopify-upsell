package database

import (
	"errors"
	"fmt"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

// ErrForeignKeyViolation is returned when an insert/update results in a foreign
// key violation
var ErrForeignKeyViolation = errors.New("not found")

// ErrNotFound is returned when a single item query results in no rows
var ErrNotFound = errors.New("not found")

// ErrUniqueViolation is returned when an insert/update results in a uniqueness violation
var ErrUniqueViolation = errors.New("already exists")

// Wrap a known error into an error that can be tested against
func Wrap(err error) error {
	if errors.Is(err, pgx.ErrNoRows) {
		return ErrNotFound
	}

	pgerror, ok := err.(*pgconn.PgError)
	if !ok {
		return err
	}
	format := func(err *pgconn.PgError) string {
		return fmt.Sprintf("%s: %s: (%s) %s", err.Severity, err.ColumnName, err.Code, err.Message)
	}
	if pgerror.Code == "23503" {
		return fmt.Errorf("%w: %s", ErrForeignKeyViolation, format(pgerror))
	}
	if pgerror.Code == "23505" {
		return fmt.Errorf("%w: %s", ErrUniqueViolation, format(pgerror))
	}
	return err
}
