package main

import "fmt"

func SomaInteiro(numeros []int) int {

	var soma = 0
	for _, numero := range numeros {
		soma += numero
	}

	return soma

}
func SomaFloat64(numeros []float64) float64 {

	var soma float64 = 0
	for _, numero := range numeros {
		soma += numero
	}

	return soma

}

func SomaAny(numeros []any) float64 {
	var soma float64 = 0
	for _, numero := range numeros {
		switch v := numero.(type) {
		case int:
			soma += float64(v)
		case float64:
			soma += v
		}
	}
	return soma
}

func SomaGenerics[T int | float64](numeros []T) T {
	var soma T = 0
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}

func main() {
	var (
		numerosInteiros = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		numerosFloat    = []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.10}
		numerosAny      = []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.10}
	)

	fmt.Println(" ### Soma sem Generics ###")
	fmt.Println(SomaInteiro(numerosInteiros))
	fmt.Println(SomaFloat64(numerosFloat))
	fmt.Println(SomaAny(numerosAny))

	fmt.Println(" ### Soma com Generics ###")
	fmt.Println(SomaGenerics(numerosInteiros))
	fmt.Println(SomaGenerics(numerosFloat))

}
