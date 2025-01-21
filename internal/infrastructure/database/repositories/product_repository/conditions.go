package product_repository

import (
	"strings"
	"wikreate/fimex/internal/domain/structure/dto/catalog_dto"
	"wikreate/fimex/internal/helpers"
)

func whereIdGroup(id_group int) (string, int) {
	return "exists(select * from categories where id = id_subcategory and id_group = ?)", id_group
}

func whereCharValues(ids []int) (string, string) {
	param := strings.Join(helpers.SliceIntValToString(ids), "','")
	return "exists(select * from product_chars where id_value in (?))", param
}

func whereChars(ids []int) (string, string) {
	param := strings.Join(helpers.SliceIntValToString(ids), "','")
	return "exists(select * from product_chars where id_char in (?))", param
}

func condGenerateNamesPayload(payload *catalog_dto.GenerateNamesInputDto) (string, []interface{}) {
	var where []string
	args := []interface{}{}

	if payload.IdGroup > 0 {
		query, param := whereIdGroup(payload.IdGroup)
		where = append(where, query)
		args = append(args, param)
	}

	if len(payload.IdsVal) > 0 {
		query, param := whereCharValues(payload.IdsVal)
		where = append(where, query)
		args = append(args, param)
	}

	if len(payload.IdsChars) > 0 {
		query, param := whereChars(payload.IdsChars)
		where = append(where, query)
		args = append(args, param)
	}

	return strings.Join(where, " and "), args
}
