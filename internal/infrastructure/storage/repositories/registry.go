package repositories

import (
	"wikreate/fimex/internal/domain/interfaces"
	"wikreate/fimex/internal/infrastructure/storage/repositories/product_repository"
	"wikreate/fimex/internal/infrastructure/storage/repositories/user_repository"
)

type Repositories struct {
	ProductRepo        *product_repository.ProductRepositoryImpl
	ProductCharRepo    *ProductCharRepositoryImpl
	UserRepo           *user_repository.UserRepositoryImpl
	PaymentHistoryRepo *PaymentHistoryRepositoryImpl
}

func NewRepositories(dbManager interfaces.DB) *Repositories {
	return &Repositories{
		ProductRepo:     product_repository.NewProductRepository(dbManager),
		ProductCharRepo: NewProductCharRepository(dbManager),

		UserRepo: user_repository.NewUserRepository(dbManager),

		PaymentHistoryRepo: NewPaymentHistoryRepository(dbManager),
	}
}
