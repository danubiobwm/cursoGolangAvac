package main

import (
	"context"
	"fmt"
	"time"

	"github.com/danubiobwm/cursoGolangAvac/cursoGolangAvac/TestesAutomatizados/Contextos/utils"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	ctx = context.WithValue(ctx, utils.UserIDKey, "user-1")

	defer cancel()

	dados, err := utils.BuscarDados(ctx, time.Second)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(dados)
}
