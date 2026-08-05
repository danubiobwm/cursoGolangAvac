package main

import (
	"sync"
	"time"
)

type Buffer struct {
	dados []int
	cond  *sync.Cond
	mu    sync.Mutex
}

func main() {
	buffer := Buffer{
		dados: make([]int, 0),
		cond:  sync.NewCond(&sync.Mutex{}),
	}

	buffer.cond = sync.NewCond(&buffer.mu)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			buffer.mu.Lock()
			println("Produzindo:", i)
			buffer.dados = append(buffer.dados, i)
			buffer.cond.Signal()
			buffer.mu.Unlock()
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			buffer.mu.Lock()
			println("Consumindo:", i)
			buffer.dados = append(buffer.dados, i)
			buffer.cond.Signal()
			buffer.mu.Unlock()
		}
	}()
	wg.Wait()
}
