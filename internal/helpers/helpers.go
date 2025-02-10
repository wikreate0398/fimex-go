package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

func IntToString(v int) string {
	return strconv.Itoa(v)
}

func TypeOf(v interface{}) reflect.Kind {
	return reflect.TypeOf(v).Kind()
}

func IsString(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.String
}

func IsStruct(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Struct
}

func ExecTime(clbk func()) {
	var start = time.Now()
	clbk()
	fmt.Println(time.Since(start))
}

func prettyprint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}

func PrintStructToJson(entity interface{}) {
	b, err := json.Marshal(entity)
	if err != nil {
		fmt.Println(err)
		return
	}
	str, _ := prettyprint(b)
	fmt.Printf("%s", str)
}
