package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		m  sync.Map
		wg sync.WaitGroup
	)

	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			m.Store(i, i)
		}()
	}

	wg.Wait()

	m.Range(func(chave, valor interface{}) bool {
		fmt.Printf("chave: %d, valor: %d\n", chave, valor)
		return true
	})

	valor, ok := m.Load(50)
	if ok {
		fmt.Printf("valor para chave 50: %d\n", valor)
	}

}
