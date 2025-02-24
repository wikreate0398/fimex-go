package db_adapter

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"wikreate/fimex/internal/domain/interfaces"
	"wikreate/fimex/pkg/mysql"
)

var _ interfaces.DB = (*DB)(nil)

type DB struct {
	db *sqlx.DB
}

func NewDBAdapter(db *sqlx.DB) *DB {
	return &DB{db: db}
}

func (dbm *DB) BeginTx() (*sqlx.Tx, error) {
	tx, err := dbm.db.Beginx()
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (dbm *DB) Get(entity interface{}, query string, args ...interface{}) error {
	if err := dbm.db.Get(entity, query, args...); err != nil {
		return err
	}
	return nil
}

func (dbm *DB) Select(entity interface{}, query string, args ...interface{}) error {
	if err := dbm.db.Select(entity, query, args...); err != nil {
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
	query, err := mysql.NewBatchUpdate(table, identifier, arg).Query()
	if err != nil {
		return nil, err
	}
	return dbm.db.MustExec(query), nil
}

func (dbm *DB) NamedExec(query string, args interface{}) error {
	if _, err := dbm.db.NamedExec(query, args); err != nil {
		return err
	}
	return nil
}
