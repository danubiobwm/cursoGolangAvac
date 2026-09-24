package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"runtime/pprof"
	"time"

	"github.com/danubiobwm/cursoGolangAvac/cursoGolangAvac/Profiling/Heap/cache"
)

func simularVazamento(c *cache.Cache) {

	for i := 0; i < 1000; i++ {
		// Adiciona itens ao cache sem removê-los
		chave := fmt.Sprintf("chave-%d", rand.IntN(1_000_000))
		tamanho := rand.IntN(100000) + 10000 // Aloca entre 10KB e 100KB
		c.Inserir(chave, tamanho)

		if len(c.Dados) > 10 {
			c.Limpar()
		}

		time.Sleep(10 * time.Millisecond) // Simula trabalho
	}
}

func main() {

	// Simula vazamento de memória
	go simularVazamento(cache.New())

	// Aguarda para capturar o perfil
	time.Sleep(3 * time.Second)

	// Gera o perfil de heap programaticamente
	f, err := os.Create("heap.prof")
	if err != nil {
		fmt.Println("Erro ao criar arquivo de perfil:", err)
		return
	}
	defer f.Close()

	if err := pprof.WriteHeapProfile(f); err != nil {
		fmt.Println("Erro ao escrever perfil de heap:", err)
		return
	}

	fmt.Println("Perfil de heap salvo em 'heap.prof'")
}
