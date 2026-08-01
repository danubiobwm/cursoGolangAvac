package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	data map[string]any
	mu   sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]any),
	}
}

func (c *Cache) Set(key string, value any) {
	logf("set: Tentativa de escrita na chave: %s", key)
	c.mu.Lock()
	defer c.mu.Unlock()
	logf("Escrevendo na chave: %s", key)
	time.Sleep(300 * time.Millisecond)
	c.data[key] = value
	logf("Escrita concluída na chave: %s", key)
}

func (c *Cache) Get(key string) (any, bool) {
	logf("get: Tentativa de leitura na chave: %s", key)
	c.mu.RLock()
	defer c.mu.RUnlock()
	logf("Lendo na chave: %s", key)
	time.Sleep(100 * time.Millisecond)
	value, ok := c.data[key]
	logf("Leitura concluída na chave: %s", key)
	return value, ok
}

var start = time.Now()

func logf(msg string, args ...any) {
	fmt.Printf("%s: %s\n", time.Since(start), fmt.Sprintf(msg, args...))
}
func main() {
	var (
		cache = NewCache()
		wg    sync.WaitGroup
	)
	wg.Add(4)

	cache.Set("chave1", "valor1")

	go func() {
		defer wg.Done()
		cache.Set("chave1", "novo valor1")
		_, _ = cache.Get("chave1")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(10 * time.Millisecond)
		cache.Set("chave2", "novo valor2")
		_, _ = cache.Get("chave2")
	}()
	go func() {
		defer wg.Done()
		time.Sleep(10 * time.Millisecond)
		cache.Set("chave3", "novo valor3")
		_, _ = cache.Get("chave3")
	}()
	go func() {
		defer wg.Done()
		time.Sleep(10 * time.Millisecond)
		cache.Set("chave4", "novo valor4")
		_, _ = cache.Get("chave4")
	}()

	wg.Wait()
}
