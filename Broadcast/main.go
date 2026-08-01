package main

import (
	"fmt"
	"time"
)

func consumidor(ch chan int, id int) {
	for {
		msg := <-ch
		fmt.Printf("Consumidor %d recebeu: %d\n", id, msg)
	}
}

func broadcast(channels []chan int, msg int) {
	for _, ch := range channels {
		ch <- msg
	}
}

func main() {
	var channels = []chan int{
		make(chan int),
		make(chan int),
		make(chan int),
	}

	for i, ch := range channels {
		//consumidor
		go consumidor(ch, i)
	}

	broadcast(channels, 142)
	time.Sleep(1 * time.Second) // Aguarda um pouco para que os consumidores recebam a mensagem
}
