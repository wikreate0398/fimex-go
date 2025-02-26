package product_service

import (
	"wikreate/fimex/internal/domain/structure/dto/catalog_dto"
)

type ProductRepository interface {
	GetIdsForGenerateNames(payload *catalog_dto.GenerateNamesInputDto, limit int, offset int) ([]string, error)
	CountTotalForGenerateNames(payload *catalog_dto.GenerateNamesInputDto) (int, error)
	GetForSort() ([]catalog_dto.ProductSortQueryDto, error)
	UpdateNames(arg interface{}, key string) error
	UpdatePosition(arg interface{}, key string) error
}

type ProductCharRepository interface {
	GetByProductIds(ids []string) ([]catalog_dto.ProductCharQueryDto, error)
}
