package mysql

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

const (
	maxSqlDBOpenConns    = 25
	maxSqlDBIdleConns    = 25
	sqlDBConnMaxLifetime = 5 * time.Minute
	timeout              = 10 * time.Second
)

type DBCreds struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func NewClient(ctx context.Context, dbConf DBCreds) (*sqlx.DB, error) {

	dsn := fmt.Sprintf(
		"%s:%s@(%s:%d)/%s",
		dbConf.User, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Database,
	)

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	db, err := sqlx.ConnectContext(ctx, "mysql", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxSqlDBOpenConns)
	db.SetMaxIdleConns(maxSqlDBIdleConns)
	db.SetConnMaxLifetime(sqlDBConnMaxLifetime)

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
