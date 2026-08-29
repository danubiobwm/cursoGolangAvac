package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type (
	Pessoa struct {
		nome  string
		idade int
	}

	Buscavel interface {
		constraints.Integer | ~string | Pessoa
	}
)

func Busca[T Buscavel](lista []T, f func(T) bool) int {
	for i, valor := range lista {
		if f(valor) {
			return i
		}
	}
	return -1
}

func main() {

	var (
		numeros = []int{1, 2, 3, 4, 5, 6}
		palavra = []string{"Ola", "mundo", "oi"}
		pessoa  = []Pessoa{
			{nome: "Davi", idade: 30},
			{nome: "Maria", idade: 40},
			{nome: "Beatriz", idade: 50},
		}
	)

	resultado := Busca(numeros, func(n int) bool {
		return n >= 3
	})

	fmt.Println("Posição do numeto que atendeu a condição: ", resultado)

	resultado2 := Busca(palavra, func(s string) bool {
		return len(s) > 3
	})

	fmt.Println("posição da string: ", resultado2)

	resultado3 := Busca(pessoa, func(p Pessoa) bool {
		return p.nome == "Maria"
	})

	fmt.Println("posição do nome: ", resultado3)

}
