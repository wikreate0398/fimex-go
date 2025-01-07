package services

import (
	"wikreate/fimex/internal/repository"
	"wikreate/fimex/internal/services/product_service"
)

type Service struct {
	ProductService *product_service.ProductService
}

func NewService(repository *repository.Repository) *Service {
	deps := &product_service.Deps{
		ProductRepository:     repository.ProductRepo,
		ProductCharRepository: repository.ProductCharRepo,
	}
	return &Service{ProductService: product_service.NewProductService(deps)}
}
