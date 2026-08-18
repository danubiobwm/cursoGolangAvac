package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("String", 1, true)
	fmt.Println(1)
	fmt.Println(true)
	fmt.Println([]int{1, 2, 3, 4})

	// --------

	var (
		numero  int = 5
		numero2 any = 5
	)

	numero = numero + 1
	numero++

	// numero2 = numero2 + 1
	// numero2++

	numero2Int, ok := numero2.(int)
	if ok {
		numero2Int++
	}

	// --------

	var (
		lista  any = []int{1, 2, 3, 4, 5, 6}
		lista2 any = []int{1, 2, 3, 4, 5, 6}
	)

	if reflect.DeepEqual(lista, lista2) {
		fmt.Println("Listas iguais!")
	} else {
		fmt.Println("Listas diferentes!")
	}

	// --------

	var x any
	x = "string"
	x = 123
	x = true

	switch v := x.(type) {
	case int:
		fmt.Println("Inteiro:", v)
	case string:
		fmt.Println("String:", v)
	case float64:
		fmt.Println("Float64:", v)
	case bool:
		fmt.Println("Booleano:", v)
	}
}
