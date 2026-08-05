package main

import (
	"fmt"
	"sync"
	"time"
)

type PoolDeConexoes struct {
	Conexoes []int
	mu       sync.Mutex
	cond     *sync.Cond
}

func NovoPoolDeConexoes(tamanho int) *PoolDeConexoes {
	pool := &PoolDeConexoes{
		Conexoes: make([]int, tamanho),
	}
	for i := 0; i < tamanho; i++ {
		pool.Conexoes[i] = i + 1
	}

	pool.cond = sync.NewCond(&pool.mu)
	return pool
}

func (p *PoolDeConexoes) ObterConexao() int {
	p.mu.Lock()
	defer p.mu.Unlock()
	for len(p.Conexoes) == 0 {
		fmt.Println("Nenhuma conexao disponivel. Aguardado...")
		p.cond.Wait()
	}
	conexao := p.Conexoes[0]
	p.Conexoes = p.Conexoes[1:]
	fmt.Printf("Conexao %d obtida\n", conexao)
	return conexao
}

func (p *PoolDeConexoes) LiberarConexao(conexao int) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.Conexoes = append(p.Conexoes, conexao)
	fmt.Printf("Conexao %d liberada\n", conexao)
	p.cond.Signal()

}

func main() {

	pool := NovoPoolDeConexoes(2)

	var wg sync.WaitGroup

	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func(id int) {
			defer wg.Done()
			conexao := pool.ObterConexao()
			fmt.Printf("Goroutine %d usando conexao %d\n", id, conexao)
			time.Sleep(time.Second * 5)
			pool.LiberarConexao(conexao)
		}(i + 1)
	}

	wg.Wait()
	fmt.Println("Todas as goroutines concluídas")
}
