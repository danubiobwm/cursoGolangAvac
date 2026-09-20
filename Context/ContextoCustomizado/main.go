package main

import (
	"context"
	"fmt"
	"time"

	"github.com/danubiobwm/cursoGolangAvac/cursoGolangAvac/Context/ContextoCustomizado/custom"
)

func main() {
	// Contexto base com timeout
	ctxPai, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Cria um contexto customizado que ignora o timeout
	customCtx := custom.NewContext(ctxPai, map[string]any{
		"chave": "valor",
	})

	// Inicia uma goroutine que verifica o comportamento do CustomContext
	go func() {
		select {
		case <-customCtx.Done(): // não será executado, pois o Done do pai nao influencia
			fmt.Println("Contexto customizado foi cancelado:", customCtx.Err())
		case <-time.After(5 * time.Second): // Simula trabalho longo
			fmt.Println("5 segundos se passaram...")
		}
	}()

	time.Sleep(6 * time.Second)
	fmt.Println("Finalizado")
}
