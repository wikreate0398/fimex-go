package interfaces

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type DbManager interface {
	GetDB() *sqlx.DB

	Get(entity interface{}, query string, args ...interface{})
	Select(entity interface{}, query string, args ...interface{})

	NamedExec(query string, args interface{})
	BatchUpdate(table string, identifier string, arg interface{}) sql.Result
}
