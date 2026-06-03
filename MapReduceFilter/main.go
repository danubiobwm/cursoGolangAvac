package main

import "fmt"

type mySlice []int

func main() {
	//map -> [1, 2, 3, 4, 5, 6, 7, 8, 9, 10] --> x2 -> [2, 4, 6, 8, 10, 12, 14, 16, 18, 20]
	//filter -> [2, 4, 6, 8, 10, 12, 14, 16, 18, 20] --> x%3 == 0 -> [6, 12, 18]
	//reduce -> [6, 12, 18] -> sum -> 36

	var lista = mySlice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	lista = lista.Filter(func(i int) bool {
		return i%2 == 0
	})
	fmt.Println("Valor da lista: ", lista)

}

func (m mySlice) Filter(cond func(int) bool) mySlice {
	var result mySlice
	for _, numero := range m {
		if cond(numero) {
			result = append(result, numero)
		}
	}
	return result
}
