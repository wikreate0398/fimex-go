package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"wikreate/fimex/pkg/database/mysql"
	"wikreate/fimex/pkg/failed"
)

type DbManager struct {
	db *sqlx.DB
}

func NewDBManager(ctx context.Context, creds mysql.DBCreds) *DbManager {
	db := mysql.NewClient(ctx, creds)
	return &DbManager{db}
}

func (dbm *DbManager) BeginTx() *sqlx.Tx {
	tx, err := dbm.db.Beginx()
	failed.PanicOnError(err, "Failed to begin transaction")
	return tx
}

func (dbm *DbManager) GetDB() *sqlx.DB {
	return dbm.db
}

func (dbm *DbManager) Get(entity interface{}, query string, args ...interface{}) {
	failed.PanicOnError(dbm.db.Get(entity, query, args...), fmt.Sprintf("Failed get query: %s", query))
}

func (dbm *DbManager) Select(entity interface{}, query string, args ...interface{}) {
	failed.PanicOnError(dbm.db.Select(entity, query, args...), fmt.Sprintf("Failed select query: %s", query))
}

func (dbm *DbManager) BatchUpdate(table string, identifier string, arg interface{}) sql.Result {
	query, err := NewBatchUpdate(table, identifier, arg).getQuery()
	failed.PanicOnError(err, fmt.Sprintf("Failed to get batch update: %s %s", table, identifier))

	//res, err := dbm.db.MustExec(query)
	//
	//if err != nil {
	//	fmt.Println(query)
	//	panic(err)
	//}

	return dbm.db.MustExec(query)
}

func (dbm *DbManager) NamedExec(query string, args interface{}) {
	_, err := dbm.db.NamedExec(query, args)
	failed.PanicOnError(err, "Failed NamedExec")
}
