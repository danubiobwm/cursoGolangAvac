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

	//map
	lista = lista.Map(func(i int) int {
		return i * 2
	})
	fmt.Println("Valor da lista: ", lista)

	//reduce
	resultado := lista.Reduce(func(acc int, i int) int {
		return acc + i
	})

	fmt.Println("Valor da lista: ", resultado)

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

func (m mySlice) Map(transform func(int) int) mySlice {
	var result mySlice
	for _, numero := range m {
		result = append(result, transform(numero))
	}
	return result
}

func (m mySlice) Reduce(acc func(int, int) int) int {
	var result int
	for _, numero := range m {
		result = acc(result, numero)
	}
	return result
}
