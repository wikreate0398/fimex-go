package catalog_dto

import (
	"database/sql"
)

type GenerateNamesInputDto struct {
	IdGroup  int   `json:"id_group"`
	IdsChars []int `json:"ids_chars"`
	IdsVal   []int `json:"ids_val"`
}

type ProductCharQueryDto struct {
	IdProduct int    `db:"id_product"`
	Name      string `db:"name"`
	UseInName bool   `db:"use_product_name"`
	UseEmoji  bool   `db:"add_emodji"`
	Position  string `db:"position"`
}

type ProductNameStoreDto struct {
	Id        int    `db:"id"`
	Name      string `db:"name"`
	UpdatedAt string `db:"updated_at"`
}

type ProductSortQueryDto struct {
	ID             int            `db:"id"`
	IdSubcategory  int            `db:"id_subcategory"`
	IdCategory     int            `db:"id_category"`
	IdBrand        int            `db:"id_brand"`
	IdGroup        int            `db:"id_group"`
	BrandPosition  int            `db:"brand_position"`
	CatPosition    int            `db:"cat_position"`
	SubCatPosition int            `db:"subcat_position"`
	Position       sql.NullString `db:"position"`
}

type ProductSortStoreDto struct {
	ID        int    `db:"id"`
	Position  int    `db:"page_up"`
	UpdatedAt string `db:"updated_at"`
}
