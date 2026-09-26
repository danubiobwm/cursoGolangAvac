package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func main() {

	// Configura o limite de threads do runtime para permitir uma análise mais interessante
	runtime.GOMAXPROCS(1)

	// Simula trabalho concorrente forçando a criação de threads
	for i := 0; i < 100; i++ {
		go criarThread(i)
	}

	// Aguarda para permitir a criação de threads
	time.Sleep(5 * time.Second)

	// Gera o perfil de threads criadas programaticamente
	f, err := os.Create("threads.prof")
	if err != nil {
		fmt.Println("Erro ao criar arquivo de perfil de threads:", err)
		return
	}
	defer f.Close()

	if err := pprof.Lookup("threadcreate").WriteTo(f, 0); err != nil {
		fmt.Println("Erro ao escrever perfil de threads:", err)
		return
	}

	fmt.Println("Perfil de threads salvo em 'threads.prof'")
}

func criarThread(id int) {
	fmt.Printf("Goroutine %d iniciada\n", id)
	time.Sleep(10 * time.Second) // Bloqueia a goroutine por um tempo
}
