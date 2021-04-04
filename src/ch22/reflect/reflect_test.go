package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int8, reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("Int")
	default:
		fmt.Println("Unknown")
	}
}

func TestCheckType(t *testing.T) {
	CheckType(11.0)
}