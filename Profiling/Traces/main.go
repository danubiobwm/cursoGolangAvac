package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

type Job struct {
	ID   int
	Tipo string
}

type Resultado struct {
	JobID   int
	Duracao time.Duration
}

func main() {
	// Cria arquivo de trace
	f, err := os.Create("trace.out")
	if err != nil {
		fmt.Println("erro ao criar trace.out:", err)
		return
	}
	defer f.Close()

	// Liga o trace
	if err := trace.Start(f); err != nil {
		fmt.Println("erro ao iniciar trace:", err)
		return
	}
	defer trace.Stop()

	// Permitir paralelismo pra ficar mais interessante no trace
	runtime.GOMAXPROCS(4)

	// Contexto com timeout: o programa roda por ~4 segundos
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	var (
		jobs         = make(chan Job, 100)
		resultados   = make(chan Resultado, 100)
		wgWorkers    sync.WaitGroup
		wgResultados sync.WaitGroup
		mu           sync.Mutex
		qtdWorkers   = 6
	)

	go produtor(ctx, jobs)

	// Alguns workers concorrentes
	wgWorkers.Add(qtdWorkers)
	for i := 0; i < qtdWorkers; i++ {
		go worker(ctx, i, jobs, resultados, &wgWorkers, &mu)
	}

	// Consumidor dos resultados
	wgResultados.Add(1)
	go agregarResultados(ctx, resultados, &wgResultados)

	// 4) Gerador de "ruído" de GC: aloca slices periodicamente
	go alocarMemoria(ctx)

	// Espera o contexto expirar
	<-ctx.Done()

	// Fecha o canal de jobs para os workers terminarem
	close(jobs)
	wgWorkers.Wait()

	// Fecha o canal de resultados e espera o sink terminar
	close(resultados)
	wgResultados.Wait()

	fmt.Println("Trace salvo em 'trace.out'")
}

// produtor gera jobs de tipos diferentes (CPU e IO) enquanto o contexto estiver ativo.
func produtor(ctx context.Context, jobs chan<- Job) {
	trace.WithRegion(ctx, "produtor", func() {
		jobID := 0
		ticker := time.NewTicker(50 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				tipo := "cpu"
				// 30% IO-bound, 70% CPU-bound só pra variar
				if rand.IntN(10) < 3 {
					tipo = "io"
				}
				j := Job{ID: jobID, Tipo: tipo}
				jobID++
				select {
				case jobs <- j:
				case <-ctx.Done():
					return
				}
			}
		}
	})
}

// worker processa jobs: alguns CPU-heavy, outros I/O-like.
// Também usa um mutex compartilhado pra criar contenção visível.
func worker(ctx context.Context, id int, jobs <-chan Job, results chan<- Resultado, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

	trace.WithRegion(ctx, fmt.Sprintf("worker-%d", id), func() {
		for {
			select {
			case <-ctx.Done():
				return
			case job, ok := <-jobs:
				if !ok {
					return
				}
				start := time.Now()

				// Seção crítica com mutex pra gerar contenção
				mu.Lock()
				simularTrabalhoCritico(id)
				mu.Unlock()

				// Trabalho principal
				switch job.Tipo {
				case "cpu":
					simularTrabalhoCPU(id)
				case "io":
					simularTrabalhoIO(id)
				}

				duracao := time.Since(start)

				select {
				case results <- Resultado{JobID: job.ID, Duracao: duracao}:
				case <-ctx.Done():
					return
				}
			}
		}
	})
}

// Pequena seção crítica fake só pra o mutex entrar no trace.
func simularTrabalhoCritico(workerID int) {
	// Não faz quase nada, só pra aparecer no trace como seção crítica.
	_ = workerID * 2
}

// Trabalho CPU-bound: faz uns loops bobos pra queimar CPU.
func simularTrabalhoCPU(workerID int) {
	// Marca essa função no trace
	trace.Logf(context.Background(), "simularTrabalhoCPU", "worker=%d", workerID)

	n := 50_000 + rand.IntN(50_000)
	soma := 0
	for i := 0; i < n; i++ {
		soma += i * i % 97
	}

	if soma%2 == 0 {
		// só pra não ser otimizado fora
		_ = soma
	}
}

// Trabalho I/O-like: simula latência externa com Sleep.
func simularTrabalhoIO(workerID int) {
	trace.Logf(context.Background(), "simularTrabalhoIO", "worker=%d", workerID)
	// Sleep entre 20ms e 80ms
	d := time.Duration(20+rand.IntN(60)) * time.Millisecond
	time.Sleep(d)
}

// agregarResultados agrega os resultados só pra manter a goroutine ocupada.
func agregarResultados(ctx context.Context, results <-chan Resultado, wg *sync.WaitGroup) {
	defer wg.Done()

	trace.WithRegion(ctx, "agregarResultados", func() {
		var count int
		var total time.Duration

		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case r, ok := <-results:
				if !ok {
					fmt.Printf("agregarResultados: processou %d resultados, tempo total %s\n", count, total)
					return
				}
				count++
				total += r.Duracao
			case <-ticker.C:
				if count > 0 {
					avg := total / time.Duration(count)
					fmt.Printf("agregarResultados: %d jobs, tempo médio %s\n", count, avg)
				}
			}
		}
	})
}

// aloca memória periodicamente pra gerar atividade de GC no trace.
func alocarMemoria(ctx context.Context) {

	trace.WithRegion(ctx, "alocarMemoria", func() {
		var dados [][]byte
		ticker := time.NewTicker(200 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				// Aloca uns ~100KB
				b := make([]byte, 100*1024)
				dados = append(dados, b)
				// De vez em quando, solta parte das referências
				if len(dados) > 20 {
					dados = dados[len(dados)/2:]
				}
			}
		}
	})
}
