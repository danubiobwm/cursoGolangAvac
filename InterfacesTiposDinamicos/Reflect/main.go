package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x any = 50
	t := reflect.TypeOf(x)

	println("Type:", t.Name())

	y := reflect.ValueOf(x)
	println("Value:", y.Int())

	if y.CanInt() {
		xInt := y.Int()
		println("Can convert to int:", xInt)
	} else {
		fmt.Println("Cannot convert to int")
	}

	fmt.Println("###")

	var y2 any = 100
	v2 := reflect.ValueOf(&y2).Elem()
	fmt.Println("Valor de y2:", v2)

	v2.Set(reflect.ValueOf(200))
	fmt.Println("Valor atual y2:", y2)
	fmt.Println("Value:", v2.Elem().Int())
}
