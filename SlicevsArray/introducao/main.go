package main

import "fmt"

func main() {
	var (
		slice = []int{1, 2, 3, 4, 5}
		array = [5]int{1, 2, 3, 4, 5}
	)
	fmt.Println("Slice:", slice)
	fmt.Println("Array:", array)

	fmt.Println("#########:")

	fmt.Println("tamanho do slice:", len(slice))
	fmt.Println("capacidade do slice:", cap(slice))

	fmt.Println("#########:")

	fmt.Println("tamanho do array:", len(array))
	fmt.Println("capacidade do array:", cap(array))

	fmt.Println("#########:")

	slice = append(slice, 6)
	fmt.Println("Slice após append:", slice)
	fmt.Println("tamanho do slice:", len(slice))
	fmt.Println("capacidade do slice:", cap(slice))

	printSlice(slice)
}

func printSlice(slice []int) {
	for i, v := range slice {
		fmt.Printf("index: %d, value: %d\n", i, v)
	}
}
