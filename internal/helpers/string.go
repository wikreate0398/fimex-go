package helpers

import "fmt"

// String
func PrependStr(str string, concat string) string {
	if str != "" {
		return fmt.Sprintf("%s %s", concat, str)
	}
	return str
}
