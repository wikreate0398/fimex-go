package repository

import (
	"wikreate/fimex/internal/domain/interfaces"
)

type Deps struct {
	DbManager interfaces.DbManager
}

type Repository struct {
	ProductRepo     *ProductRepositoryImpl
	ProductCharRepo *ProductCharRepositoryImpl
}

func NewRepository(dbManager interfaces.DbManager) *Repository {
	deps := &Deps{dbManager}

	return &Repository{
		ProductRepo:     NewProductRepository(deps),
		ProductCharRepo: NewProductCharRepository(deps),
	}
}
