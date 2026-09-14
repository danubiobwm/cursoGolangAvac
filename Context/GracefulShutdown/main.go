// Criar Servidor HTTP com duas rotas
// - Uma rota simplesmente retorna um OK
// - A outra rota retorna um OK depois de 5 segundos

// Graceful shutdown:
// deve esperar todas as requests finalizarem antes de encerrar a aplicação.
// Porem, se mais de 10 segundos se passarem e as requests não tiverem sido finalizadas, o encerramento deve ser forçado.

package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	e.GET("/slow", func(c echo.Context) error {
		time.Sleep(5 * time.Second)
		return c.String(http.StatusOK, "Retorou depois de 5 segundos")
	})

	e.GET("/fast", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	fmt.Println("Servidor inicializado com Sucesso!")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	fmt.Println("Recebendo sinal de interrupção...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal("Erro no graceful shutdown: %s", err.Error())
	}

	fmt.Println("Servidor encerrado com sucesso!")
}
