package consumers

import (
	"wikreate/fimex/internal/dto/app_dto"
)

type Consumers struct {
	GenerateProductsNamesQueue *GenerateNamesConsumer
	SortProductsQueue          *SortProductsConsumer
	RecalcBallanceHistory      *RecalcHistoryBallanceConsumer
}

func NewHandlers(application *app_dto.Application) *Consumers {
	productService := application.Deps.Services.ProductService
	paymentHistoryService := application.Deps.Services.PaymentHistoryService

	return &Consumers{
		GenerateProductsNamesQueue: NewGenerateProductsNamesConsumer(productService),
		SortProductsQueue:          NewSortProductsConsumer(productService),

		RecalcBallanceHistory: NewRecalcHistoryBallanceConsumer(paymentHistoryService),
	}
}
