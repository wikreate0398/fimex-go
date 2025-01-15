package interfaces

import (
	"wikreate/fimex/internal/domain/entities/catalog/product_entities"
	"wikreate/fimex/internal/domain/structure/inputs"
)

type ProductRepository interface {
	GetIdsForGenerateNames(payload *inputs.GenerateNamesPayloadInput, limit int, offset int) []string
	CountTotalForGenerateNames(payload *inputs.GenerateNamesPayloadInput) int
	CountTotal() int
	GetForSort() []product_entities.ProductSortDto
	UpdateNames(arg interface{}, key string)
	UpdatePosition(arg interface{}, key string)
}

type ProductCharRepository interface {
	GetByProductIds(ids []string) []product_entities.ProductCharDto
}
