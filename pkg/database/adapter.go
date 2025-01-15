package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"wikreate/fimex/pkg/database/mysql"
)

type Logger interface {
	Panic(args ...interface{})
	Fatal(args ...interface{})
}

type DbAdapter struct {
	db     *sqlx.DB
	logger Logger
}

func NewDBManager(ctx context.Context, creds mysql.DBCreds, logger Logger) *DbAdapter {
	db, err := mysql.NewClient(ctx, creds)
	if err != nil {
		logger.Fatal(err, "Failed to connect to database")
	}
	return &DbAdapter{db: db, logger: logger}
}

func (dbm *DbAdapter) BeginTx() *sqlx.Tx {
	tx, err := dbm.db.Beginx()
	if err != nil {
		dbm.logger.Panic(err, "Failed to begin transaction")
	}
	return tx
}

func (dbm *DbAdapter) GetDB() *sqlx.DB {
	return dbm.db
}

func (dbm *DbAdapter) Get(entity interface{}, query string, args ...interface{}) {
	err := dbm.db.Get(entity, query, args...)
	if err != nil {
		dbm.logger.Panic(err, "Failed get query", map[string]any{"query": query, "args": args})
	}
}

func (dbm *DbAdapter) Select(entity interface{}, query string, args ...interface{}) {
	err := dbm.db.Select(entity, query, args...)
	if err != nil {
		dbm.logger.Panic(err, "Failed select query", map[string]any{"query": strings.ReplaceAll(query, "\\n", " "), "args": args})
	}
}

func (dbm *DbAdapter) BatchUpdate(table string, identifier string, arg interface{}) sql.Result {
	query, err := NewBatchUpdate(table, identifier, arg).query()
	if err != nil {
		dbm.logger.Panic(err, fmt.Sprintf("Failed to get batch update: %s %s", table, identifier), nil)
	}
	return dbm.db.MustExec(query)
}

func (dbm *DbAdapter) NamedExec(query string, args interface{}) {
	_, err := dbm.db.NamedExec(query, args)
	if err != nil {
		dbm.logger.Panic(err,
			"Failed NamedExec",
			map[string]any{"query": query, "args": args},
		)
	}
}
