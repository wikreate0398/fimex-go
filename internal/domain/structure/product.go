package structure

type ProductChar struct {
	IdProduct int    `db:"id_product"`
	Name      string `db:"name"`
	UseInName bool   `db:"use_product_name"`
	UseEmoji  bool   `db:"add_emodji"`
	Position  string `db:"position"`
}
