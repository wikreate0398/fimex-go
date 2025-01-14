package helpers

func SliceIntToString(slice []int) []string {
	var stringSlice []string
	for i := 0; i < len(slice); i++ {
		stringSlice = append(stringSlice, ToString(slice[i]))
	}
	return stringSlice
}
