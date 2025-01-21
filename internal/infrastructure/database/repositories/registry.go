package repositories

import (
	"wikreate/fimex/internal/dto/repo_dto"
	"wikreate/fimex/internal/infrastructure/database/repositories/product_repository"
	"wikreate/fimex/internal/infrastructure/database/repositories/user_repository"
	"wikreate/fimex/pkg/database"
)

type Repositories struct {
	ProductRepo     *product_repository.ProductRepositoryImpl
	ProductCharRepo *ProductCharRepositoryImpl
	UserRepo        *user_repository.UserRepositoryImpl
}

func NewRepositories(dbManager *database.DbAdapter) *Repositories {
	deps := &repo_dto.Deps{DbManager: dbManager}

	return &Repositories{
		ProductRepo:     product_repository.NewProductRepository(deps),
		ProductCharRepo: NewProductCharRepository(deps),

		UserRepo: user_repository.NewUserRepository(deps),
	}
}
