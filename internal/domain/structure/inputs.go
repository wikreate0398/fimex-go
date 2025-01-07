package structure

type GenerateNamesPayloadInput struct {
	IdGroup  int            `json:"id_group"`
	IdsChars map[int]string `json:"ids_chars"`
	IdsVal   map[int]string `json:"ids_val"`
}
