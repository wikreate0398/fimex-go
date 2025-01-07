package consumers

import (
	"encoding/json"
	"wikreate/fimex/internal/domain/structure"
	"wikreate/fimex/internal/services/product_service"
)

type GenerateNamesConsumer struct {
	service *product_service.ProductService
	input   *structure.GenerateNamesPayloadInput
}

func NewGenerateNamesConsumer(service *product_service.ProductService) Consumer {
	return &GenerateNamesConsumer{service, nil}
}

func (r *GenerateNamesConsumer) Handle() {
	r.service.GenerateNames(r.input)
}

func (r *GenerateNamesConsumer) ToStruct(result []byte) {
	json.Unmarshal(result, &r.input)
}
