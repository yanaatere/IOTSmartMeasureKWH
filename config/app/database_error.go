package app

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"github.com/lib/pq"
)

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	ErrUniqueViolation     = Error("unique_violation")
	ErrNullValueNotAllowed = Error("null_value_not_allowed")
	ErrorUndefinedTable    = Error("undefined_table")
	ErrNoRowsFound         = Error("no rows found")
)

const canceledMessage = "pq: canceling statement due to transaction request"

func ParsePostgresSQLError(err error) error {
	switch err {
	case sql.ErrNoRows:
		return ErrNoRowsFound
	case driver.ErrBadConn:
		return context.DeadlineExceeded
	}

	switch et := err.(type) {
	case *pq.Error:
		switch et.Code {
		case "23505":
			return ErrUniqueViolation
		case "42P01":
			return ErrorUndefinedTable
		case "22004":
			return ErrNullValueNotAllowed
		}
	}

	return err
}
