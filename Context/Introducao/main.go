package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, "chave", "valor 22")

	leitura(ctx)

}

func leitura(ctx context.Context) {
	valorArmazenado := ctx.Value("chave")

	if valorArmazenado != nil {
		fmt.Println("Valor: ", valorArmazenado)
	} else {
		fmt.Println("Valor não encontrado")
	}
}
