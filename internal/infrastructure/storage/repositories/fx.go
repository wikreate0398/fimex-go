package repositories

import (
	"go.uber.org/fx"
	"wikreate/fimex/internal/infrastructure/storage/repositories/product_repository"
	"wikreate/fimex/internal/infrastructure/storage/repositories/user_repository"
)

var Module = fx.Module("repositories",
	fx.Provide(NewPaymentHistoryRepository),
	fx.Provide(NewProductCharRepository),
	fx.Provide(user_repository.NewUserRepository),
	fx.Provide(product_repository.NewProductRepository),
)
