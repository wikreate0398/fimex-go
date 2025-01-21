package product

import (
	"wikreate/fimex/internal/domain/structure/dto/catalog_dto"
	"wikreate/fimex/internal/infrastructure/adapters/emoji"
)

type ProductChar struct {
	idProduct int
	name      string
	useInName bool
	useEmoji  bool
	position  string
}

func NewProductChar(dto catalog_dto.ProductCharQueryDto) *ProductChar {
	return &ProductChar{
		idProduct: dto.IdProduct,
		name:      dto.Name,
		useInName: dto.UseInName,
		useEmoji:  dto.UseEmoji,
		position:  dto.Position,
	}
}

func (pc *ProductChar) PrepareNameForProduct() string {
	if !pc.useInName {
		return ""
	}

	if pc.useEmoji {
		str := emoji.FindAll(pc.name)
		if len(str) > 0 {
			return str[0].Character
		}
		return pc.name
	}

	return emoji.RemoveEmojis(pc.name)
}
