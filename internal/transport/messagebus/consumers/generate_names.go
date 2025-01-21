package consumers

import (
	"encoding/json"
	"wikreate/fimex/internal/domain/services/catalog/product_service"
	"wikreate/fimex/internal/domain/structure/dto/catalog_dto"
)

type GenerateNamesConsumer struct {
	service *product_service.ProductService
}

func NewGenerateProductsNamesConsumer(
	service *product_service.ProductService,
) *GenerateNamesConsumer {
	return &GenerateNamesConsumer{service}
}

func (r *GenerateNamesConsumer) Handle(result []byte) error {
	var input = new(catalog_dto.GenerateNamesInputDto)
	if err := json.Unmarshal(result, &input); err != nil {
		return err
	}

	r.service.GenerateNames(input)

	return nil
}
