package main

import (
	"fmt"
	"reflect"
)

func main() {
	var floatVar float64 = 1.123
	v := reflect.ValueOf(floatVar)
	vp := reflect.ValueOf(&floatVar)
	fmt.Printf("Float Addressable: %v\n", v.CanAddr())
	fmt.Printf("Float Canset: %v\n", v.CanSet())
	fmt.Printf("FloatP Addressable: %v\n", vp.Elem().CanAddr())
	fmt.Printf("FloatP Canset: %v\n", vp.Elem().CanSet())
	vp.Elem().SetFloat(3.1415)
	fmt.Println("Float Changed: ", floatVar)
}
