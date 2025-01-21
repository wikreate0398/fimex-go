package consumers

import (
	"wikreate/fimex/internal/domain/services/catalog/product_service"
)

type SortProductsConsumer struct {
	service *product_service.ProductService
}

func NewSortProductsConsumer(service *product_service.ProductService) *SortProductsConsumer {
	return &SortProductsConsumer{service}
}

func (r *SortProductsConsumer) Handle(result []byte) error {
	r.service.Sort()
	return nil
}
