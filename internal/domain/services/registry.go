package services

import (
	"wikreate/fimex/internal/domain/services/catalog/product_service"
	"wikreate/fimex/internal/infrastructure/database/repositories"
)

type Services struct {
	ProductService *product_service.ProductService
}

func NewServices(repository *repositories.Repositories) *Services {
	deps := &product_service.Deps{
		ProductRepository:     repository.ProductRepo,
		ProductCharRepository: repository.ProductCharRepo,
	}
	return &Services{ProductService: product_service.NewProductService(deps)}
}
