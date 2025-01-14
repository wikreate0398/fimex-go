package repositories

import (
	"wikreate/fimex/internal/domain/interfaces"
	"wikreate/fimex/internal/domain/structure/dto/repo_dto"
	"wikreate/fimex/internal/infrastructure/database/repositories/product_repository"
)

type Repositories struct {
	ProductRepo     *product_repository.ProductRepositoryImpl
	ProductCharRepo *ProductCharRepositoryImpl
}

func NewRepositories(dbManager interfaces.DbManager) *Repositories {
	deps := &repo_dto.Deps{DbManager: dbManager}

	return &Repositories{
		ProductRepo:     product_repository.NewProductRepository(deps),
		ProductCharRepo: NewProductCharRepository(deps),
	}
}
