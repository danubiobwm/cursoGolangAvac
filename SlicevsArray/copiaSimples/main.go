package main

import "fmt"

func main() {

	fmt.Println("Cópia Simples de Slices em Go / insegura para modificações")
	// Criando um slice de inteiros
	sliceOriginal := []int{1, 2, 3, 4, 5}
	var slice2 = sliceOriginal // Isso não cria uma cópia, ambos apontam para o mesmo array subjacente

	fmt.Println("Slice Original:", sliceOriginal)
	fmt.Println("Slice 2:", slice2)

	slice2[0] = 10 // Modificando o primeiro elemento do slice2

	fmt.Println("Slice Original após modificação:", sliceOriginal)
	fmt.Println("Slice 2 após modificação:", slice2)

	slice2 = append(slice2, 6) // Adicionando um novo elemento ao slice2
	slice2[0] = 99             // Modificando o primeiro elemento do slice2 novamente
	fmt.Println("Slice Original após append:", sliceOriginal)
	fmt.Println("Slice 2 após append:", slice2)
	fmt.Println("###########")

	fmt.Println("Cópia Segura de Slices em Go / usando make e copy")
	// Criando um slice de inteiros
	var slice3 = []int{1, 2, 3, 4, 5}
	// Criando uma cópia segura do slice3
	var slice4 = make([]int, len(slice3))
	copy(slice4, slice3)

	fmt.Println("Slice 3:", slice3)
	fmt.Println("Slice 4:", slice4)

	slice4[0] = 10 // Modificando o primeiro elemento do slice4

	fmt.Println("Slice 3 após modificação:", slice3)
	fmt.Println("Slice 4 após modificação:", slice4)
}
