package main

import (
	"fmt"
	"sync"
)

type Cache struct {
	data map[string]any
	mu   sync.Mutex
}

func newCache() *Cache {
	return &Cache{
		data: make(map[string]any),
	}
}

func (c *Cache) Set(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
	fmt.Printf("Set key: %s, value: %v\n", key, value)
}

func (c *Cache) Get(key string) (any, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, ok := c.data[key]
	return value, ok
}

func main() {

	var (
		cache = newCache()
		wg    sync.WaitGroup
	)

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			cache.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))

		}()
	}

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			value, ok := cache.Get(key)
			if ok {
				fmt.Printf("Get key: %s, value: %v\n", key, value)
			} else {
				fmt.Printf("Get key: %s, value not found\n", key)
			}
		}()
	}

	wg.Wait()

}
