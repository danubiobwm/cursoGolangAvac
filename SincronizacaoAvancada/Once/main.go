package main

import (
	"fmt"
	"sync"
	"time"
)

type BancoDados struct {
	conectado bool
}

var (
	bancoDados *BancoDados
	once       sync.Once
)

func ObterBancoDados() *BancoDados {
	once.Do(func() {
		fmt.Println("Conectando ao banco de dados...")
		bancoDados = &BancoDados{conectado: true}
		time.Sleep(time.Second * 5)
		fmt.Println("Conexão estabelecida.")
	})
	return bancoDados
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func(id int) {
			defer wg.Done()
			bd := ObterBancoDados()
			fmt.Printf("Goroutine %d: Banco de dados conectado? %v\n", id, bd.conectado)
		}(i)
	}
	wg.Wait()

	bd := ObterBancoDados()
	fmt.Printf("Banco de dados conectado? %v\n", bd.conectado)
	fmt.Println("Programa finalizado.")
}
