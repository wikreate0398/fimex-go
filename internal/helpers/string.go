package helpers

import "fmt"

func PrependStr(str string, concat string) string {
	if str != "" {
		return fmt.Sprintf("%s %s", concat, str)
	}
	return str
}
