package repository

import (
	"fmt"
	"strings"
	"wikreate/fimex/internal/domain/structure"
)

type ProductCharRepositoryImpl struct {
	deps *Deps
}

func NewProductCharRepository(deps *Deps) *ProductCharRepositoryImpl {
	return &ProductCharRepositoryImpl{deps}
}

func (p ProductCharRepositoryImpl) GetByProductIds(ids []string) []structure.ProductChar {
	var productChars []structure.ProductChar

	query := fmt.Sprintf(`
			select id_product,name,use_product_name,add_emodji,cgc.position 
			from product_chars as pc
			join chars on chars.id = pc.id_value 
			join catalog_groups_chars as cgc on cgc.id_char = pc.id_char and cgc.id_group = pc.id_group
			where id_product in (%s) 
			and use_product_name = 1 
			and chars.deleted_at is null
			order by chars.page_up 
		`, strings.Join(ids, ","))

	p.deps.DbManager.Select(&productChars, query)

	return productChars
}
