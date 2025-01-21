package services

import (
	"wikreate/fimex/internal/domain/services/catalog/product_service"
	"wikreate/fimex/internal/domain/services/payment_history_service"
	"wikreate/fimex/internal/infrastructure/database/repositories"
)

type Services struct {
	ProductService        *product_service.ProductService
	PaymentHistoryService *payment_history_service.PaymentHistoryService
}

func NewServices(repository *repositories.Repositories) *Services {
	return &Services{
		ProductService: product_service.NewProductService(&product_service.Deps{
			ProductRepository:     repository.ProductRepo,
			ProductCharRepository: repository.ProductCharRepo,
		}),

		PaymentHistoryService: payment_history_service.NewPaymentHistoryService(&payment_history_service.Deps{
			UserRepo: repository.UserRepo,
		}),
	}
}
