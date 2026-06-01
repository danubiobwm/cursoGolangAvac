package main

import "fmt"

type Pessoa struct {
	Nome  *string
	Idade int
}

func main() {
	var nome = "Pedro"

	pessoa1 := Pessoa{
		Nome:  &nome,
		Idade: 26,
	}

	pessoa2 := pessoa1

	fmt.Println("Nome da pessoa 1:", *pessoa1.Nome)
	fmt.Println("Nome da pessoa 2:", *pessoa2.Nome)

	pessoa2.Nome = toPointer("Davi")

	fmt.Println("Nome da pessoa 1:", *pessoa1.Nome)
	fmt.Println("Nome da pessoa 2:", *pessoa2.Nome)

	fmt.Println("--------")

	pessoa3 := Pessoa{
		Nome:  toPointer("Leonardo"),
		Idade: 26,
	}

	pessoa4 := deepCopy(pessoa3)

	*pessoa4.Nome = "Lucas"
	fmt.Println("Nome da pessoa 3:", *pessoa3.Nome)
	fmt.Println("Nome da pessoa 4:", *pessoa4.Nome)

	fmt.Println("--------")
	listaDePessoas1 := []Pessoa{pessoa1, pessoa2, pessoa3, pessoa4}
	listaDePessoas2 := deepCopyList(listaDePessoas1)

	fmt.Println("Lista 1", listaDePessoas1)
	fmt.Println("Lista 2", listaDePessoas2)
}

func toPointer(s string) *string {
	return &s
}

func deepCopy(origem Pessoa) Pessoa {
	var destino Pessoa

	destino.Idade = origem.Idade
	destino.Nome = toPointer(*origem.Nome)

	return destino
}

func deepCopyList(origem []Pessoa) []Pessoa {
	var destino = make([]Pessoa, len(origem))

	for i, pessoa := range origem {
		destino[i] = deepCopy(pessoa)
	}

	return destino
}
