package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	fmt.Println("Type of: ", reflect.TypeOf(x))
	fmt.Println("Value: ", reflect.ValueOf(x).String())
	// From value, you can find type, value, kind of types etc.
	v := reflect.ValueOf(x)
	fmt.Println("Type: ", v.Type())
	fmt.Println("Kind of float 64: ", v.Kind() == reflect.Float64)
	fmt.Println("Content: ", v.Float())
	fmt.Printf("Value from relect to Interface Value: %7.1e\n", v.Interface())
}
