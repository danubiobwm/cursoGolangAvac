package main

func etapa1(ch1 chan<- int, numeros []int) {
	for _, n := range numeros {
		ch1 <- n
	}

	close(ch1)
}
func etapa2(ch1 <-chan int, ch2 chan<- int) {
	for numero := range ch1 {
		ch2 <- numero * numero
	}
	close(ch2)
}

func etapa3(ch2 <-chan int) {
	for numero := range ch2 {
		println("Valor recebido na etapa 3", numero)
	}
}

func main() {
	var (
		numeros = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		ch1     = make(chan int)
		ch2     = make(chan int)
	)

	go etapa1(ch1, numeros)
	go etapa2(ch1, ch2)
	etapa3(ch2)

}
