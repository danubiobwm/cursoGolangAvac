package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func main() {
	// Simula alocações intensivas
	for i := 0; i < 5; i++ {
		simularAlocacoesIntensivas(i)
	}

	// Gera o perfil de Allocs
	f, err := os.Create("allocs.prof")
	if err != nil {
		fmt.Println("Erro ao criar arquivo de perfil de allocs:", err)
		return
	}
	defer f.Close()

	profile := pprof.Lookup("allocs")
	if err := profile.WriteTo(f, 0); err != nil {
		fmt.Println("Erro ao escrever perfil de allocs:", err)
		return
	}

	fmt.Println("Perfil de allocs salvo em 'allocs.prof'")
}

func simularAlocacoesIntensivas(iteracao int) {
	fmt.Printf("Processando iteração %d\n", iteracao)
	dados := make([][]byte, 1000)

	// Simula alocações de memória
	for i := 0; i < 1000; i++ {
		dados[i] = make([]byte, 1024*10) // Aloca 10KB por item
	}

	// Simula processamento (mantendo parte dos dados na memória)
	time.Sleep(500 * time.Millisecond)

	// Intencionalmente não limpa todos os dados
	for i := 0; i < 500; i++ {
		dados[i] = nil // Libera parte dos dados
	}

	runtime.GC() // Força o garbage collector para coletar memória liberada
}
