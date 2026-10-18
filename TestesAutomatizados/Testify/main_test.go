package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContarPalavras(t *testing.T) {
	input := "Olá mundo olá go go go"
	outputEsperado := map[string]int{"olá": 2, "mundo": 1, "go": 3}

	output := ContarPalavras(input)

	assert.Equal(
		t,
		outputEsperado,
		output,
		"ContarPalavras(%q) retornou um mapa inesperado", input,
	)

}
