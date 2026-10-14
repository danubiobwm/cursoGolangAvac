package main

import (
	"fmt"
	"strings"
)

func main() {
	palavras := ContarPalavras("Olá mundo olá go go go")

	fmt.Println(palavras)
}

func ContarPalavras(s string) map[string]int {
	var (
		palavras       = strings.Fields(strings.ToLower(s))
		mapaDePalavras = make(map[string]int)
	)

	for _, palavra := range palavras {
		mapaDePalavras[palavra]++
	}

	return mapaDePalavras
}
