package product_service

import (
	"wikreate/fimex/internal/domain/structure/dto/catalog_dto"
)

type ProductRepository interface {
	GetIdsForGenerateNames(payload *catalog_dto.GenerateNamesInputDto, limit int, offset int) []string
	CountTotalForGenerateNames(payload *catalog_dto.GenerateNamesInputDto) int
	CountTotal() int
	GetForSort() []catalog_dto.ProductSortQueryDto
	UpdateNames(arg interface{}, key string)
	UpdatePosition(arg interface{}, key string)
}

type ProductCharRepository interface {
	GetByProductIds(ids []string) []catalog_dto.ProductCharQueryDto
}
