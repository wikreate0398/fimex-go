package database

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"wikreate/fimex/pkg/database/mysql"
)

type DB struct {
	db *sqlx.DB
}

func NewMysqlManager(ctx context.Context, creds mysql.DBCreds) (*DB, error) {
	db, err := mysql.NewClient(ctx, creds)

	if err != nil {
		return nil, err
	}

	return &DB{db: db}, nil
}

func (dbm *DB) BeginTx() (*sqlx.Tx, error) {
	tx, err := dbm.db.Beginx()
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (dbm *DB) GetDB() *sqlx.DB {
	return dbm.db
}

func (dbm *DB) Get(entity interface{}, query string, args ...interface{}) error {
	err := dbm.db.Get(entity, query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (dbm *DB) Select(entity interface{}, query string, args ...interface{}) error {
	err := dbm.db.Select(entity, query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (dbm *DB) Query(query string, args ...any) (*sql.Rows, error) {
	rows, err := dbm.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (dbm *DB) BatchUpdate(table string, identifier string, arg interface{}) (sql.Result, error) {
	query, err := NewBatchUpdate(table, identifier, arg).query()
	if err != nil {
		return nil, err
	}
	return dbm.db.MustExec(query), nil
}

func (dbm *DB) NamedExec(query string, args interface{}) error {
	_, err := dbm.db.NamedExec(query, args)
	if err != nil {
		return err
	}
	return nil
}
