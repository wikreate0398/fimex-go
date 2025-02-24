package consumers

import (
	"wikreate/fimex/internal/domain/structure/dto/catalog_dto"
	"wikreate/fimex/internal/domain/structure/dto/payment_dto"
)

type PaymentHistoryService interface {
	RecalcBallances(payload *payment_dto.RecalcBallanceInputDto)
}

type ProductService interface {
	GenerateNames(payload *catalog_dto.GenerateNamesInputDto)
	Sort()
}
