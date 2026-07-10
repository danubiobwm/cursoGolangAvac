package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d iniciado\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d terminado\n", id)
}
func workerErrorGroup(id int) error {
	fmt.Printf("Worker error group %d iniciado\n", id)
	if id == 2 {
		return fmt.Errorf("Worker error group %d encontrou um erro", id)
	}
	time.Sleep(time.Second)
	fmt.Printf("Worker error group %d terminado\n", id)
	return nil
}

func main() {
	fmt.Println("--- WaitGroup ----")

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("Todos os workers terminado")

	fmt.Println("--- ErrorGroup ----")
	var eg errgroup.Group

	for i := 0; i < 3; i++ {
		i := i // Capturamos as variáveis para evitar problemas de concorrência
		eg.Go(func() error {
			return workerErrorGroup(i)
		})
	}
	err := eg.Wait()
	if err != nil {
		fmt.Printf("Error workers: %v\n", err)
	} else {
		fmt.Println("Todos os workers terminaram sem erros")
	}
}
