package interfaces

import "wikreate/fimex/internal/domain/structure"

type ProductRepository interface {
	GetIds(payload *structure.GenerateNamesPayloadInput, limit int, offset int) []string
	UpdateNames(arg interface{}, key string)
	CountTotal(payload *structure.GenerateNamesPayloadInput) int
}

type ProductCharRepository interface {
	GetByProductIds(ids []string) []structure.ProductChar
}
