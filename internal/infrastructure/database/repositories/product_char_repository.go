package repositories

import (
	"fmt"
	"strings"
	"wikreate/fimex/internal/domain/interfaces"
	"wikreate/fimex/internal/domain/structure/dto/catalog_dto"
)

type ProductCharRepositoryImpl struct {
	dbManager interfaces.DbManager
}

func NewProductCharRepository(db interfaces.DbManager) *ProductCharRepositoryImpl {
	return &ProductCharRepositoryImpl{dbManager: db}
}

func (p ProductCharRepositoryImpl) GetByProductIds(ids []string) []catalog_dto.ProductCharQueryDto {
	var productChars []catalog_dto.ProductCharQueryDto
	query := fmt.Sprintf(`
			select id_product,name,use_product_name,add_emodji,cgc.position 
			from product_chars as pc
			join chars on chars.id = pc.id_value 
			join catalog_groups_chars as cgc on cgc.id_char = pc.id_char and cgc.id_group = pc.id_group
			where id_product in (%s) 
			and use_product_name = 1 
			and chars.deleted_at is null  
		`, strings.Join(ids, ","))

	p.dbManager.Select(&productChars, query)

	return productChars
}
