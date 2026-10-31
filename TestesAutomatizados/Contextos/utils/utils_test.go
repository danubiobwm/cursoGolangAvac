package utils

import (
	"context"
	"testing"
	"time"

	"github.com/danubiobwm/cursoGolangAvac/cursoGolangAvac/TestesAutomatizados/Contextos/erros"
	"github.com/stretchr/testify/assert"
)

func TestBuscarDados(t *testing.T) {
	t.Parallel()

	testes := []struct {
		nome           string
		setupCtx       func() context.Context
		delay          time.Duration
		outputEsperado string
		erroEsperado   error
	}{
		{
			nome: "sucesso com timeout e userID no contexto",
			setupCtx: func() context.Context {
				ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
				t.Cleanup(cancel) // garante cancel no fim do teste
				return context.WithValue(ctx, UserIDKey, "user-123")
			},
			delay:          50 * time.Millisecond,
			outputEsperado: "dados fictícios recuperados para o usuário user-123",
			erroEsperado:   nil,
		},
		{
			nome: "cancelamento antes de terminar",
			setupCtx: func() context.Context {
				ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
				ctx = context.WithValue(ctx, UserIDKey, "user-456")

				// cancela depois de um tempinho
				go func() {
					time.Sleep(20 * time.Millisecond)
					cancel()
				}()

				return ctx
			},
			delay:          100 * time.Millisecond,
			outputEsperado: "",
			erroEsperado:   context.Canceled,
		},
		{
			nome: "timeout estourando antes do delay",
			setupCtx: func() context.Context {
				ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
				t.Cleanup(cancel)
				return context.WithValue(ctx, UserIDKey, "user-789")
			},
			delay:          100 * time.Millisecond,
			outputEsperado: "",
			erroEsperado:   context.DeadlineExceeded,
		},
		{
			nome: "erro quando nao ha userID no contexto",
			setupCtx: func() context.Context {
				// tem deadline, mas nao tem userID
				ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
				t.Cleanup(cancel)
				return ctx
			},
			delay:          50 * time.Millisecond,
			outputEsperado: "",
			erroEsperado:   erros.ErroUserIDObrigatorio,
		},
		{
			nome: "erro quando contexto nao tem deadline",
			setupCtx: func() context.Context {
				// tem userID, mas nao tem deadline
				return context.WithValue(context.Background(), UserIDKey, "user-sem-deadline")
			},
			delay:          50 * time.Millisecond,
			outputEsperado: "",
			erroEsperado:   erros.ErroDeadlineObrigatorio,
		},
	}

	for _, teste := range testes {
		t.Run(teste.nome, func(t *testing.T) {
			t.Parallel()

			result, err := BuscarDados(teste.setupCtx(), teste.delay)

			if teste.erroEsperado == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.ErrorIs(t, err, teste.erroEsperado)
			}

			assert.Equal(t, teste.outputEsperado, result)
		})
	}
}
