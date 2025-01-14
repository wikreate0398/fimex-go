package inputs

type GenerateNamesPayloadInput struct {
	IdGroup  int   `json:"id_group"`
	IdsChars []int `json:"ids_chars"`
	IdsVal   []int `json:"ids_val"`
}
