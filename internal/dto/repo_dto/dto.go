package repo_dto

import (
	"wikreate/fimex/pkg/database"
)

type Deps struct {
	DbManager *database.DbAdapter
}
