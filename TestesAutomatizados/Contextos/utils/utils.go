package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/danubiobwm/cursoGolangAvac/cursoGolangAvac/TestesAutomatizados/Contextos/erros"
)

// chave de contexto com tipo próprio, pra evitar colisão
type ctxKey string

const UserIDKey ctxKey = "userID"

func BuscarDados(ctx context.Context, delay time.Duration) (string, error) {

	// 1) Valida valor obrigatório no contexto
	userID, ok := ctx.Value(UserIDKey).(string)
	if !ok || userID == "" {
		return "", erros.ErroUserIDObrigatorio
	}

	// 2) Valida que o contexto tem deadline configurado
	if _, ok := ctx.Deadline(); !ok {
		return "", erros.ErroDeadlineObrigatorio
	}

	// 3) Executa "trabalho" respeitando cancel/timeout
	select {
	case <-time.After(delay): // simula serviço demorando
		return fmt.Sprintf("dados fictícios recuperados para o usuário %s", userID), nil
	case <-ctx.Done(): // cancel / timeout
		return "", ctx.Err()
	}
}
