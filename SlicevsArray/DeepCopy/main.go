package main

import "fmt"

func main() {
	// usando matriz
	original := [5]int{1, 2, 3, 4, 5}
	deepCopy := original

	fmt.Println("Original:", original)
	fmt.Println("Deep Copy:", deepCopy)

	matrix := [5][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10}}
	fmt.Println("Original Matrix:", matrix)
	deepCopyMatrixArr := matrix
	fmt.Println("Deep Copy Matrix:", deepCopyMatrixArr)

	//// Outro exemplo usando
	fmt.Println("Outro exemplo usando")

	matrix1 := [][]int{
		{1, 2},
		{3, 4}}
	matrix2 := make([][]int, len(matrix1))

	copy(matrix2, matrix1)
	fmt.Println("Original Matrix:", matrix1)
	fmt.Println("Deep Copy Matrix:", matrix2)

	matrix2[0] = []int{9, 9}
	fmt.Println("Original Matrix after modification:", matrix1)
	fmt.Println("Deep Copy Matrix after modification:", matrix2)

	fmt.Println("Usando função de deep copy")
	matrix3 := [][]int{
		{1, 2},
		{3, 4}}
	matrix4 := deepCopyMatrix(matrix3)
	matrix4[0] = []int{9, 9}
	fmt.Println("Original Matrix:", matrix3)
	fmt.Println("Deep Copy Matrix:", matrix4)
}

func deepCopyMatrix(matrix [][]int) [][]int {

	destino := make([][]int, len(matrix))

	for i, slice := range matrix {
		destino[i] = make([]int, len(slice))
		copy(destino[i], slice)
	}

	return destino
}
