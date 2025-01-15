package product_entities

import "database/sql"

type ProductCharDto struct {
	IdProduct int    `db:"id_product"`
	Name      string `db:"name"`
	UseInName bool   `db:"use_product_name"`
	UseEmoji  bool   `db:"add_emodji"`
	Position  string `db:"position"`
}

func (p ProductCharDto) ToEntity() *ProductChar {
	return NewProductCharEntity(p)
}

type ProductNameDto struct {
	Id        int    `db:"id"`
	Name      string `db:"name"`
	UpdatedAt string `db:"updated_at"`
}

type ProductSortDto struct {
	ID             int            `db:"id"`
	IdSubcategory  int            `db:"id_subcategory"`
	IdCategory     int            `db:"id_category"`
	IdBrand        int            `db:"id_brand"`
	IdGroup        int            `db:"id_group"`
	BrandPosition  string         `db:"brand_position"`
	CatPosition    string         `db:"cat_position"`
	SubCatPosition string         `db:"subcat_position"`
	PageUp         int            `db:"page_up"`
	Position       sql.NullString `db:"position"`
}

type ProductInsertSortDto struct {
	ID        int    `db:"id"`
	Position  int    `db:"page_up"`
	UpdatedAt string `db:"updated_at"`
}
