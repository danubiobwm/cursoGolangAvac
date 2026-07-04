package main

import (
	"fmt"
	"time"
)

func enviar(ch chan int) {
	fmt.Println("Enviado o valor para o chanel")
	ch <- 100
	fmt.Println("Valor Enviado com sucesso")
}

func receber(ch chan int) {
	value := <-ch
	fmt.Println("Valor recebido", value)
}

func main() {
	// Create a channel of type int
	// Channels are used to communicate between goroutines
	// The make function is used to create a channel
	ch := make(chan int)

	for i := 0; i < 3; i++ {
		go enviar(ch) // chanel sem buffer, ou seja, o envio e recebimento precisam estar sincronizados
	}

	time.Sleep(time.Second)

	for i := 0; i < 3; i++ {
		go receber(ch)
	}

	time.Sleep(time.Second * 5)

}
