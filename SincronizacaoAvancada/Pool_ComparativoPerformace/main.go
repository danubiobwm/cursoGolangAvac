package main

import (
	"bytes"
	"fmt"
	"sync"
	"time"
)

func semPool(iteracoes int) time.Duration {
	start := time.Now()
	for i := 0; i < iteracoes; i++ {
		buffer := bytes.Buffer{}
		buffer.WriteString("Testando")
		_ = buffer.String()
	}
	return time.Since(start)
}
func comPool(iteracoes int) time.Duration {
	start := time.Now()

	pool := sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
	for i := 0; i < iteracoes; i++ {
		buffer := pool.Get().(*bytes.Buffer)
		buffer.Reset()
		buffer.WriteString("Testando")
		_ = buffer.String()
		buffer.Reset()
		pool.Put(buffer)
	}
	return time.Since(start)
}
func porcentagem(t1, t2 time.Duration) float64 {
	return float64(t1-t2) / float64(t1) * 100
}

func main() {
	iteracoes := 1_000_000

	fmt.Printf("Executado comparação com %d iterações\n", iteracoes)
	tl := semPool(iteracoes)
	fmt.Printf("Sem pool tempo: %s\n", tl)
	t2 := comPool(iteracoes)
	fmt.Printf("Com pool tempo: %s\n", t2)
	p := porcentagem(tl, t2)
	fmt.Printf("Percentual de melhoria: %.2f%%\n", p)

}
