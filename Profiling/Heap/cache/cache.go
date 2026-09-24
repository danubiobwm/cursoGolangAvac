package cache

import "fmt"

type Cache struct {
	Dados map[string][]byte
}

func New() *Cache {
	return &Cache{
		Dados: make(map[string][]byte),
	}
}

func (c *Cache) Inserir(chave string, tamanho int) {
	c.Dados[chave] = make([]byte, tamanho)
	fmt.Printf("Adicionado: %s (tamanho: %d bytes)\n", chave, tamanho)
}

func (c *Cache) Limpar() {
	for k := range c.Dados {
		delete(c.Dados, k)
		fmt.Printf("Removido: %s\n", k)
		return
	}
}
