package main

import "fmt"

func main() {
	// estrutura de chave e valor
	// chave: valor

	// var meuMap = map[string]any{
	// 	"nome":      "Gustavo",
	// 	"sobrenome": "Gomes",
	// }
	// fmt.Println(meuMap)
	// fmt.Println(meuMap["nome"])
	// fmt.Println(meuMap["sobrenome"])

	var meuMap = map[string]any{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
	}

	for chave, valor := range meuMap {
		fmt.Printf("A chave é %s e o valor é %v\n", chave, valor)
	}

	fmt.Println("######")

	var meuMap2 = map[string]any{
		"nome":      "Gustavo",
		"sobrenome": "Gomes",
	}

	for chave, valor := range meuMap2 {
		fmt.Printf("A chave é %s e o valor é %v\n", chave, valor)
	}
}
