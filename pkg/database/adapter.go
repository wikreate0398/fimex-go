package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"wikreate/fimex/pkg/database/mysql"
	"wikreate/fimex/pkg/failed"
)

type DbAdapter struct {
	db *sqlx.DB
}

func NewDBManager(ctx context.Context, creds mysql.DBCreds) *DbAdapter {
	db := mysql.NewClient(ctx, creds)
	return &DbAdapter{db}
}

func (dbm *DbAdapter) BeginTx() *sqlx.Tx {
	tx, err := dbm.db.Beginx()
	failed.PanicOnError(err, "Failed to begin transaction")
	return tx
}

func (dbm *DbAdapter) GetDB() *sqlx.DB {
	return dbm.db
}

func (dbm *DbAdapter) Get(entity interface{}, query string, args ...interface{}) {
	failed.PanicOnError(dbm.db.Get(entity, query, args...), fmt.Sprintf("Failed get query: %s", query))
}

func (dbm *DbAdapter) Select(entity interface{}, query string, args ...interface{}) {
	failed.PanicOnError(dbm.db.Select(entity, query, args...), fmt.Sprintf("Failed select query: %s", query))
}

func (dbm *DbAdapter) BatchUpdate(table string, identifier string, arg interface{}) sql.Result {
	query, err := NewBatchUpdate(table, identifier, arg).getQuery()
	failed.PanicOnError(err, fmt.Sprintf("Failed to get batch update: %s %s", table, identifier))
	return dbm.db.MustExec(query)
}

func (dbm *DbAdapter) NamedExec(query string, args interface{}) {
	_, err := dbm.db.NamedExec(query, args)
	failed.PanicOnError(err, "Failed NamedExec")
}
