package main

import (
	"fmt"
	"time"
)

func agendar(intervalo time.Duration, f func(), chanelDeParada <-chan struct{}, chanelDeFinalização chan<- struct{}) {
	ticket := time.NewTicker(intervalo)
	defer ticket.Stop()
	defer close(chanelDeFinalização)

	for {
		select {
		case <-ticket.C:
			f()
		case <-chanelDeParada:
			fmt.Println("Parando a função agendada")
			return
		}
	}
}

func main() {
	var chanelDeParada = make(chan struct{})      // stop
	var chanelDeFinalização = make(chan struct{}) // done

	go agendar(3*time.Second, func() {
		println("Executando a função agendada")
	}, chanelDeParada, chanelDeFinalização)

	time.Sleep(10 * time.Second)

	close(chanelDeParada)
	<-chanelDeFinalização
}
