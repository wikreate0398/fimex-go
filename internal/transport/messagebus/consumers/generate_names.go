package consumers

import (
	"encoding/json"
	"wikreate/fimex/internal/domain/interfaces"
	"wikreate/fimex/internal/domain/services/catalog/product_service"
	"wikreate/fimex/internal/domain/structure/inputs"
)

type GenerateNamesConsumer struct {
	service *product_service.ProductService
	input   *inputs.GenerateNamesPayloadInput
	log     interfaces.Logger
}

func NewGenerateProductsNamesConsumer(
	service *product_service.ProductService, log interfaces.Logger,
) *GenerateNamesConsumer {
	return &GenerateNamesConsumer{service, nil, log}
}

func (r *GenerateNamesConsumer) Handle() {
	r.service.GenerateNames(r.input)
}

func (r *GenerateNamesConsumer) ToStruct(result []byte) {
	err := json.Unmarshal(result, &r.input)
	r.log.PanicOnErr(err, "Unmarshal failed")
}
