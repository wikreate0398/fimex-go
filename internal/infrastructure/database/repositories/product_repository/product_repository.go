package product_repository

import (
	"fmt"
	"wikreate/fimex/internal/domain/structure/dto/catalog_dto"
	"wikreate/fimex/internal/dto/repo_dto"
	"wikreate/fimex/internal/helpers"
)

type ProductRepositoryImpl struct {
	deps *repo_dto.Deps
}

func NewProductRepository(deps *repo_dto.Deps) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{deps}
}

func (p ProductRepositoryImpl) GetIdsForGenerateNames(payload *catalog_dto.GenerateNamesInputDto, limit int, offset int) []string {
	var cond, args = condGenerateNamesPayload(payload)
	args = append(args, limit, offset)

	var ids []string

	var query = fmt.Sprintf("select id from products %s order by id asc LIMIT ? OFFSET ?", helpers.PrependStr(cond, "where"))

	p.deps.DbManager.Select(&ids, query, args...)

	return ids
}

func (p ProductRepositoryImpl) CountTotalForGenerateNames(payload *catalog_dto.GenerateNamesInputDto) int {
	var cond, args = condGenerateNamesPayload(payload)

	var total int
	var query = fmt.Sprintf("select count(*) from products %s", helpers.PrependStr(cond, "where"))

	p.deps.DbManager.Get(&total, query, args...)
	return total
}

func (p ProductRepositoryImpl) CountTotal() int {
	var total int
	var query = "select count(*) from products"

	p.deps.DbManager.Get(&total, query)
	return total
}

func (p ProductRepositoryImpl) GetForSort() []catalog_dto.ProductSortQueryDto {
	query := `
		SELECT 
			products.id, 
			products.id_subcategory, 
			products.id_category, 
			categories.id_brand,   
			categories.id_group, 
			products.page_up, 
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
		INNER JOIN categories AS categories ON categories.id = products.id_category
		LEFT JOIN categories AS subcategory ON subcategory.id = products.id_subcategory
		INNER JOIN brands ON brands.id = categories.id_brand
		WHERE products.deleted_at IS NULL`

	var dto []catalog_dto.ProductSortQueryDto
	p.deps.DbManager.Select(&dto, query)
	return dto
}

func (p ProductRepositoryImpl) UpdateNames(arg interface{}, identifier string) {
	p.deps.DbManager.BatchUpdate("products", identifier, arg)
}

func (p ProductRepositoryImpl) UpdatePosition(arg interface{}, identifier string) {
	p.deps.DbManager.BatchUpdate("products", identifier, arg)
}
