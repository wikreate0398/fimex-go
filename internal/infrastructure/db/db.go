package db

import (
	"context"
	"wikreate/fimex/internal/config"
	"wikreate/fimex/internal/domain/interfaces"
	"wikreate/fimex/internal/infrastructure/adapters/db_adapter"
	"wikreate/fimex/pkg/mysql"
)

func NewDb(cfg *config.Config) (interfaces.DB, error) {
	conf := cfg.Databases.MySql
	ctx := context.Background()

	db, err := mysql.NewClient(ctx, mysql.DBCreds{
		Host:     conf.Host,
		Port:     conf.Port,
		User:     conf.User,
		Password: conf.Password,
		Database: conf.Database,
	})

	if err != nil {
		return nil, err
	}

	return db_adapter.NewDBAdapter(db), nil
}
