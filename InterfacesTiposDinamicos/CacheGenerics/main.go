package main

import "fmt"

type (
	Cache[K comparable, V any] struct {
		data map[K]V
	}
)

func NovoCache[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{
		data: make(map[K]V),
	}
}

func (c *Cache[K, V]) Set(key K, value V) {
	c.data[key] = value
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	value, ok := c.data[key]
	return value, ok
}

func main() {
	cache := NovoCache[string, int]()
	cache.data["um"] = 1
	cache.data["dois"] = 2

	cache.Set("tres", 3)
	cache.Set("quatro", 4)

	fmt.Println(cache.data["um"])   // Output: 1
	fmt.Println(cache.data["dois"]) // Output: 2

	fmt.Println(cache.Get("um"))     // Output: 1 true
	fmt.Println(cache.Get("dois"))   // Output: 2 true
	fmt.Println(cache.Get("tres"))   // Output: 3 true
	fmt.Println(cache.Get("quatro")) // Output: 4 true

}
