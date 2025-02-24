package interfaces

import (
	"database/sql"
)

type DB interface {
	Get(entity interface{}, query string, args ...interface{}) error
	Select(entity interface{}, query string, args ...interface{}) error
	Query(query string, args ...any) (*sql.Rows, error)

	NamedExec(query string, args interface{}) error
	BatchUpdate(table string, identifier string, arg interface{}) (sql.Result, error)
}
