package consumers

import (
	"wikreate/fimex/internal/domain/structure/dto/app_dto"
)

type Consumers struct {
	GenerateProductsNamesQueue *GenerateNamesConsumer
	SortProductsQueue          *SortProductsConsumer
}

func NewHandlers(application *app_dto.Application) *Consumers {
	productService := application.Deps.Services.ProductService

	return &Consumers{
		GenerateProductsNamesQueue: NewGenerateProductsNamesConsumer(productService, application.Deps.Logger),
		SortProductsQueue:          NewSortProductsConsumer(productService),
	}
}
