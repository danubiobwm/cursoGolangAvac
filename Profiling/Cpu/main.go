package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"

	"math/rand/v2"
)

func main() {
	// Criação do arquivo de perfil (profile) de CPU
	f, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Println("Erro ao criar arquivo de perfil:", err)
		return
	}
	defer f.Close()

	// Início do profiling de CPU
	if err := pprof.StartCPUProfile(f); err != nil {
		fmt.Println("Erro ao iniciar perfil de CPU:", err)
		return
	}
	defer pprof.StopCPUProfile()

	// Executa uma carga de trabalho complexa
	trabalhoComplexo()

	fmt.Println("Perfil de CPU salvo em 'cpu.prof'")
}

func trabalhoComplexo() {
	resultados := make(chan int, 5)

	// Executa operações concorrentes
	for i := 0; i < 5; i++ {
		go func() {
			resultados <- somaPesada()
		}()
	}

	// Aguarda os resultados
	for i := 0; i < 5; i++ {
		resultado := <-resultados
		fmt.Printf("Resultado %d: %d\n", i+1, resultado)
	}

	// Realiza operações sequenciais
	for i := 0; i < 3; i++ {
		calculoPesado()
		fmt.Printf("Cálculo pesado realizado. Iteração: %d\n", i)
	}
}

func somaPesada() int {
	total := 0
	for i := 0; i < 1_000_000; i++ {
		total += rand.IntN(100)
	}
	return total
}

func calculoPesado() {
	time.Sleep(500 * time.Millisecond) // simula I/O
	for i := 0; i < 500_000; i++ {
		_ = i * i
	}
}
