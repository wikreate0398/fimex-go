package helpers

import (
	"fmt"
	"reflect"
	"time"
)

func ToString(v interface{}) string {
	return reflect.ValueOf(v).String()
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
