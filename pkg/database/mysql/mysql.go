package mysql

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
	"wikreate/fimex/pkg/failed"
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

func NewClient(ctx context.Context, dbConf DBCreds) *sqlx.DB {

	dsn := fmt.Sprintf(
		"%s:%s@(%s:%d)/%s",
		dbConf.User, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Database,
	)

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	
	db, err := sqlx.ConnectContext(ctx, "mysql", dsn)
	failed.PanicOnError(err, "Failed to connect to the database")

	db.SetMaxOpenConns(maxSqlDBOpenConns)
	db.SetMaxIdleConns(maxSqlDBIdleConns)
	db.SetConnMaxLifetime(sqlDBConnMaxLifetime)

	err = db.PingContext(ctx)
	failed.PanicOnError(err, "Failed to ping the database")

	return db
}
