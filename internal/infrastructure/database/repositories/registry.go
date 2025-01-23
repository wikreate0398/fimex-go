package repositories

import (
	"wikreate/fimex/internal/domain/interfaces"
	"wikreate/fimex/internal/infrastructure/database/repositories/product_repository"
	"wikreate/fimex/internal/infrastructure/database/repositories/user_repository"
)

type Repositories struct {
	ProductRepo        *product_repository.ProductRepositoryImpl
	ProductCharRepo    *ProductCharRepositoryImpl
	UserRepo           *user_repository.UserRepositoryImpl
	PaymentHistoryRepo *PaymentHistoryRepositoryImpl
}

func NewRepositories(dbManager interfaces.DbManager) *Repositories {
	return &Repositories{
		ProductRepo:     product_repository.NewProductRepository(dbManager),
		ProductCharRepo: NewProductCharRepository(dbManager),

		UserRepo: user_repository.NewUserRepository(dbManager),

		PaymentHistoryRepo: NewPaymentHistoryRepository(dbManager),
	}
}
