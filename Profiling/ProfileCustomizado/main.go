package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"sync"
	"time"
)

// Perfil customizado que conta "tarefas ativas".
// A métrica aqui é: quantos itens ainda estão ativos (Add sem Remove) por stack trace.
var tarefasAtivas = pprof.NewProfile("tarefas-ativas")

// tarefa é só um marcador sem dados; o pprof só liga a "tarefa" à stack de criação.
type tarefa struct {
	id int
}

func main() {
	var wg sync.WaitGroup

	// Sobe algumas tarefas LENTAS
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go tarefaLenta(&wg, i)
	}

	// Sobe várias tarefas RÁPIDAS
	for i := 5; i < 25; i++ {
		wg.Add(1)
		go tarefaRapida(&wg, i)
	}

	// Espera um pouco antes de tirar o snapshot.
	// Nesse momento, é bem provável que:
	// - quase todas as tarefas rápidas já tenham terminado
	// - várias tarefas devagares ainda estejam ativas
	time.Sleep(500 * time.Millisecond)

	f, err := os.Create("tarefas-ativas.prof")
	if err != nil {
		fmt.Println("erro ao criar arquivo:", err)
		return
	}
	defer f.Close()

	if err := pprof.Lookup("tarefas-ativas").WriteTo(f, 0); err != nil {
		// Exporta o profile customizado
		fmt.Println("erro ao exportar profile:", err)
		return
	}

	fmt.Println("Perfil 'tarefas-ativas' salvo em tarefas-ativas.prof")

	// Espera todas as goroutines terminarem antes de sair
	wg.Wait()
}

// Tarefa rápida: fica pouco tempo ativa
func tarefaRapida(wg *sync.WaitGroup, id int) {
	defer wg.Done()

	t := &tarefa{id: id}
	// Marca essa tarefa como "ativa" no profile
	tarefasAtivas.Add(t, 0)
	defer tarefasAtivas.Remove(t)

	// Simula pouco trabalho
	time.Sleep(100 * time.Millisecond)
}

// Tarefa lenta: fica mais tempo ativa
func tarefaLenta(wg *sync.WaitGroup, id int) {
	defer wg.Done()

	t := &tarefa{id: id}
	tarefasAtivas.Add(t, 0)
	defer tarefasAtivas.Remove(t)

	// Simula trabalho mais longo
	time.Sleep(1 * time.Second)
}
