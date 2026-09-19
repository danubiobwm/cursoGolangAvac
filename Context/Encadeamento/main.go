package main

import (
	"context"
	"fmt"
	"time"
)

type myString string

const (
	timestamp myString = "timestamp"
	id        myString = "id"
)

func escreverCtx(identificador string, ctx context.Context) {
	fmt.Printf("[%-6s] timestamp: %v, id=%v\n", identificador, ctx.Value(timestamp), ctx.Value(id))
}

func main() {

	ctxPai, cancelPai := context.WithCancel(
		context.WithValue(context.Background(), timestamp, time.Now().Format(time.RFC3339)),
	)

	ctxFilho1, cancelFilho1 := context.WithCancel(
		context.WithValue(ctxPai, id, "context-filho1"),
	)
	ctxFilho2, _ := context.WithCancel(
		context.WithValue(ctxPai, id, "context-filho2"),
	)

	fmt.Println("== Valores dentro do context ==")
	escreverCtx("Pai", ctxPai)
	escreverCtx("Filho1", ctxFilho1)
	escreverCtx("Filho2", ctxFilho2)

	go func() {

		<-ctxPai.Done()
		fmt.Println("Contexto pai foi cancelado", ctxPai.Err())

	}()
	go func() {

		<-ctxFilho1.Done()
		fmt.Println("Contexto filho1 foi cancelado", ctxFilho1.Err())

	}()

	go func() {

		<-ctxFilho2.Done()
		fmt.Println("Contexto filho2 foi cancelado", ctxFilho2.Err())

	}()

	time.Sleep(time.Second * 3)
	fmt.Println("Cancelando o contexto filho1")
	cancelFilho1()

	time.Sleep(time.Second * 3)
	fmt.Println("Cancelando o contexto Pai")
	cancelPai()

	time.Sleep(time.Second * 3)

}
