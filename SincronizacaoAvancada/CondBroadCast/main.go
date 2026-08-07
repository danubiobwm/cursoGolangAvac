package main

import (
	"fmt"
	"sync"
	"time"
)

type Trava struct {
	liberada bool
	mu       sync.Mutex
	cond     *sync.Cond
}

func NovaTrava() *Trava {
	t := &Trava{}
	t.cond = sync.NewCond(&t.mu)
	return t
}

var start = time.Now()

func logf(msg string) {
	fmt.Printf("[%4dm] %s\n", time.Since(start).Microseconds(), msg)
}

func worker(id int, t *Trava, wg *sync.WaitGroup) {
	defer wg.Done()
	logf(fmt.Sprintf("Worker %d: aguardando liberação da trava", id))
	t.mu.Lock()
	defer t.mu.Unlock()

	for !t.liberada {
		logf(fmt.Sprintf("Worker %d: trava não liberada, aguardando...", id))
		t.cond.Wait()
	}
	logf(fmt.Sprintf("Worker %d: trava liberada, continuando execução", id))
	time.Sleep(time.Millisecond * 100)
}

func main() {

	{ // bloco de escopo para limitar a vida útil das variáveis
		logf("Main 1: iniciando workers e aguardando liberação da trava")
		var (
			qtdeWorkers = 5
			trava       = NovaTrava()
			wg          sync.WaitGroup
		)

		wg.Add(qtdeWorkers)

		for i := 0; i < qtdeWorkers; i++ {
			go worker(i, trava, &wg)
		}
		time.Sleep(time.Microsecond * 500)
		trava.mu.Lock()
		trava.liberada = true
		logf("Main: liberando trava e notificando todos os workers")
		trava.cond.Broadcast()
		trava.mu.Unlock()
		wg.Wait()
	}
	{
		logf("Signal 2: iniciando workers e aguardando liberação da trava")
		var (
			qtdeWorkers = 4
			trava       = NovaTrava()
			wg          sync.WaitGroup
		)

		wg.Add(qtdeWorkers)

		for i := 0; i < qtdeWorkers; i++ {
			go worker(i, trava, &wg)
		}
		time.Sleep(time.Microsecond * 500)
		trava.mu.Lock()
		trava.liberada = true
		logf("Signal: liberando trava e notificando um worker")
		trava.cond.Signal()
		trava.mu.Unlock()

		time.Sleep(time.Microsecond * 300)
		trava.mu.Lock()
		logf("Signal: call signal again to notify another worker")
		trava.cond.Signal()
		trava.mu.Unlock()

		time.Sleep(time.Microsecond * 200)
		trava.mu.Lock()
		logf("Signal: call signal again to notify another worker")
		trava.cond.Signal()
		trava.mu.Unlock()

		time.Sleep(time.Microsecond * 100)
		trava.mu.Lock()
		logf("Signal: call signal again to notify another worker")
		trava.cond.Signal()
		trava.mu.Unlock()

		wg.Wait()
	}

}
