package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"net/http/pprof"
	"time"
)

func main() {
	// Mux só para pprof, isolado na porta 6060
	pprofMux := http.NewServeMux()

	pprofMux.HandleFunc("/debug/pprof/", pprof.Index)          // Página índice: lista e serve todos os profiles (/heap, /goroutine, /mutex, etc.)
	pprofMux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline) // Mostra a linha de comando usada pra iniciar o processo (só info extra pro pprof)
	pprofMux.HandleFunc("/debug/pprof/profile", pprof.Profile) // Gera perfil de CPU por alguns segundos (usado por `go tool pprof` em /profile)
	pprofMux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)   // Endpoint que o pprof usa pra traduzir endereços de memória em nomes de funções/arquivos
	pprofMux.HandleFunc("/debug/pprof/trace", pprof.Trace)     // Gera trace de execução (usado por `go tool trace` em /trace)

	// Sobe o servidor de profiling em outra goroutine
	go func() {
		log.Println("pprof ouvindo em :6060")
		if err := http.ListenAndServe(":6060", pprofMux); err != nil {
			log.Println("erro no servidor de pprof:", err)
		}
	}()

	// ========= Servidor principal (porta 8080) =========

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá! Tente chamar /work pra gerar carga de CPU :)")
	})

	http.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		result := trabalhoPesado()
		duracao := time.Since(start)

		fmt.Fprintf(w, "Trabalho pesado concluído!\n")
		fmt.Fprintf(w, "Resultado: %d\n", result)
		fmt.Fprintf(w, "Demorou: %s\n", duracao)
	})

	log.Println("app ouvindo em :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// trabalhoPesado simula uma carga de CPU pra aparecer no profile
func trabalhoPesado() int {
	total := 0

	for i := 0; i < 30_000_000; i++ {
		total += rand.IntN(100)
	}

	return total
}
