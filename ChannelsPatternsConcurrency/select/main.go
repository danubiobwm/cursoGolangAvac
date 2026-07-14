package main

import (
	"fmt"
	"time"
)

func enviarMensagem(ch chan<- string, mensagem string, intervalo time.Duration) {
	for {
		time.Sleep(intervalo)
		ch <- mensagem
	}
}

func main() {

	var (
		ch1 = make(chan string)
		ch2 = make(chan string)
	)

	go enviarMensagem(ch1, "Mensagem do canal 1", 2*time.Second)
	go enviarMensagem(ch2, "Mensagem do canal 2", 3*time.Second)

	fmt.Println("Aguardando mensagens dos canais...")

	for i := 0; i < 10; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Recebido do canal 1:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Recebido do canal 2:", msg2)
		case <-time.After(5 * time.Second):
			fmt.Println("Nenhuma mensagem recebida nos últimos 5 segundos.")
			return
		}
	}

}
