package main

import (
	"fmt"
	"sync"
)

func produtor(numeros []int, entrada chan<- int) {
	for _, n := range numeros {
		entrada <- n
	}
	close(entrada)
}

func fanOut(id int, entrada <-chan int, saida chan<- int) {
	for n := range entrada {
		fmt.Printf("Worker %d processando %d\n", id, n)
		saida <- n * n
	}
	close(saida)
}

func fanIn(saidaReadOnly []<-chan int) <-chan int {
	var (
		saidaFinal = make(chan int)
		wg         sync.WaitGroup
	)

	wg.Add(len(saidaReadOnly))

	for _, saida := range saidaReadOnly {
		go func() {
			defer wg.Done()
			for n := range saida {
				saidaFinal <- n
			}
		}()
	}
	go func() {
		wg.Wait()
		close(saidaFinal)
	}()
	return saidaFinal
}

func main() {
	var (
		numeros = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //simulação de workload
		entrada = make(chan int)
		saidas  = []chan int{make(chan int), make(chan int), make(chan int)} //fanout
	)

	go produtor(numeros, entrada)

	for i, saida := range saidas {
		go fanOut(i+1, entrada, saida)
		fmt.Printf("Worker %d iniciado\n", i)
	}

	saidaReadOnly := make([]<-chan int, len(saidas))

	for i, saida := range saidas {
		saidaReadOnly[i] = saida
	}

	saidaFinal := fanIn(saidaReadOnly)

	for resultado := range saidaFinal {
		fmt.Printf("Resultado final: %d\n", resultado)
	}

}
