package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		m  = make(map[int]int)
		mu sync.Mutex
		wg sync.WaitGroup
	)

	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			m[i] = i
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	for chave, valor := range m {
		fmt.Printf("chave: %d, valor: %d\n", chave, valor)
	}

}
