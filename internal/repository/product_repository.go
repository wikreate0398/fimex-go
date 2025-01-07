package repository

import (
	"fmt"
	"wikreate/fimex/internal/domain/structure"
)

type ProductRepositoryImpl struct {
	deps *Deps
}

func NewProductRepository(deps *Deps) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{deps}
}

func (p ProductRepositoryImpl) GetIds(payload *structure.GenerateNamesPayloadInput, limit int, offset int) []string {
	var where string
	if payload.IdGroup > 0 {
		where += fmt.Sprintf(`
			where exists(select * from categories where id = id_subcategory and id_group = %v)
		`, payload.IdGroup)
	}

	var ids []string
	p.deps.DbManager.Select(&ids, "select id from products LIMIT ? OFFSET ?", limit, offset)

	return ids
}

func (p ProductRepositoryImpl) CountTotal(payload *structure.GenerateNamesPayloadInput) int {
	var where string
	if payload.IdGroup > 0 {
		where += fmt.Sprintf(`
			where exists(select * from categories where id = id_subcategory and id_group = %v)
		`, payload.IdGroup)
	}

	var total int
	p.deps.DbManager.Get(&total, "select count(*) from products")
	return total
}

func (p ProductRepositoryImpl) UpdateNames(arg interface{}, identifier string) {
	p.deps.DbManager.BatchUpdate("products", identifier, arg)
}
