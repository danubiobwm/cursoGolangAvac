package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"sync"
	"time"
)

func main() {
	var (
		wg       sync.WaitGroup
		bloqueio = make(chan bool)
	)

	// Simula várias goroutines concorrentes
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d iniciada\n", id)
			time.Sleep(1 * time.Second) // Simula trabalho
		}(i)
	}

	// Simula goroutines bloqueadas
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d bloqueada\n", id+100)
			<-bloqueio // Bloqueia indefinidamente
		}(i)
	}

	// Aguarda algumas goroutines finalizarem antes de capturar o perfil
	time.Sleep(2 * time.Second)

	// Gera o perfil de goroutines programaticamente
	f, err := os.Create("goroutines.prof")
	if err != nil {
		fmt.Println("Erro ao criar arquivo de perfil:", err)
		return
	}
	defer f.Close()

	if err := pprof.Lookup("goroutine").WriteTo(f, 0); err != nil {
		fmt.Println("Erro ao escrever perfil de goroutines:", err)
		return
	}

	fmt.Println("Perfil de goroutines salvo em 'goroutines.prof'")

	// Finaliza as goroutines restantes
	close(bloqueio)
	wg.Wait()
}
