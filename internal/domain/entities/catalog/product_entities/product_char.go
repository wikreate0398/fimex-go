package product_entities

import (
	"wikreate/fimex/internal/adapters/emoji"
)

type ProductChar struct {
	idProduct int
	name      string
	useInName bool
	useEmoji  bool
	position  string
}

func NewProductCharEntity(dto ProductCharDto) *ProductChar {
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
