package main

import (
	"fmt"
	"sync"
)

type Cache struct {
	data sync.Map
}

func newCache() *Cache {
	return &Cache{}
}

func (c *Cache) Set(key string, value any) {
	c.data.Store(key, value)
	fmt.Printf("Set key: %s, value: %v\n", key, value)
}

func (c *Cache) Get(key string) (any, bool) {
	return c.data.Load(key)
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
