package repositories

import (
	"go.uber.org/fx"
	"wikreate/fimex/internal/domain/services/catalog/product_service"
	"wikreate/fimex/internal/domain/services/payment_history_service"
	"wikreate/fimex/internal/infrastructure/storage/repositories/product_repository"
	"wikreate/fimex/internal/infrastructure/storage/repositories/user_repository"
)

var _ product_service.ProductCharRepository = (*ProductCharRepositoryImpl)(nil)
var _ product_service.ProductRepository = (*product_repository.ProductRepositoryImpl)(nil)

var Module = fx.Module("repositories",
	fx.Provide(
		fx.Annotate(
			NewPaymentHistoryRepository,
			fx.As(new(payment_history_service.PaymentHistoryRepository)),
		),

		fx.Annotate(
			user_repository.NewUserRepository,
			fx.As(new(payment_history_service.UserRepository)),
		),

		fx.Annotate(
			NewProductCharRepository,
			fx.As(new(product_service.ProductCharRepository)),
		),

		fx.Annotate(
			product_repository.NewProductRepository,
			fx.As(new(product_service.ProductRepository)),
		),
	),
)
