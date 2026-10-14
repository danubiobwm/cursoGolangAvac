package main

import (
	"reflect"
	"testing"
)

func TestContarPalavras(t *testing.T) {
	input := "Olá mundo olá go go go"

	outputEsperado := map[string]int{"olá": 2, "mundo": 1, "go": 3} // want
	output := ContarPalavras(input)                                 // got

	if !reflect.DeepEqual(output, outputEsperado) {
		t.Errorf("ContarPalavras(%v) = %v; esperado %v", input, output, outputEsperado)
	}
}
