package product_repository

import (
	"fmt"
	"wikreate/fimex/internal/domain/interfaces"
	"wikreate/fimex/internal/domain/structure/dto/catalog_dto"
	"wikreate/fimex/internal/helpers"
)

type ProductRepositoryImpl struct {
	db interfaces.DB
}

func NewProductRepository(db interfaces.DB) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{db: db}
}

func (p ProductRepositoryImpl) GetIdsForGenerateNames(payload *catalog_dto.GenerateNamesInputDto, limit int, offset int) ([]string, error) {

	var cond, args = condGenerateNamesPayload(payload)
	args = append(args, limit, offset)

	var ids []string

	var query = fmt.Sprintf("select id from products %s order by id asc LIMIT ? OFFSET ?", helpers.PrependStr(cond, "where"))

	if err := p.db.Select(&ids, query, args...); err != nil {
		return nil, err
	}

	return ids, nil
}

func (p ProductRepositoryImpl) CountTotalForGenerateNames(payload *catalog_dto.GenerateNamesInputDto) (int, error) {
	var cond, args = condGenerateNamesPayload(payload)

	var total int
	var query = fmt.Sprintf("select count(*) from products %s", helpers.PrependStr(cond, "where"))

	if err := p.db.Get(&total, query, args...); err != nil {
		return 0, err
	}

	return total, nil
}

func (p ProductRepositoryImpl) GetForSort() ([]catalog_dto.ProductSortQueryDto, error) {
	query := `
		SELECT 
			products.id, 
			products.id_subcategory, 
			products.id_category, 
			categories.id_brand,   
			categories.id_group,  
			brands.page_up as brand_position,
			categories.page_up as cat_position,
			subcategory.page_up as subcat_position,
			(
				SELECT GROUP_CONCAT(
					chars.page_up
					ORDER BY cgc.position SEPARATOR ','
				)
				FROM chars
				JOIN product_chars AS pc ON pc.id_value = chars.id AND pc.id_product = products.id   
				JOIN catalog_groups_chars AS cgc ON cgc.id_char = pc.id_char AND cgc.id_group = pc.id_group
				WHERE cgc.in_bot = 1 
				AND pc.id_product = products.id 
				AND exclude = 0
			) AS position
		FROM products
		JOIN categories AS categories ON categories.id = products.id_category
		JOIN categories AS subcategory ON subcategory.id = products.id_subcategory
		JOIN brands ON brands.id = categories.id_brand
		WHERE products.deleted_at IS NULL  
		order by brand_position, cat_position, subcat_position`

	var dto []catalog_dto.ProductSortQueryDto

	if err := p.db.Select(&dto, query); err != nil {
		return nil, err
	}

	return dto, nil
}

func (p ProductRepositoryImpl) UpdateNames(arg interface{}, identifier string) error {
	_, err := p.db.BatchUpdate("products", identifier, arg)
	return err
}

func (p ProductRepositoryImpl) UpdatePosition(arg interface{}, identifier string) error {
	_, err := p.db.BatchUpdate("products", identifier, arg)
	return err
}
