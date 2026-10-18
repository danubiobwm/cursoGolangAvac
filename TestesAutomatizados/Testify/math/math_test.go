package math

import (
	"reflect"
	"testing"

	"github.com/danubiobwm/cursoGolangAvac/cursoGolangAvac/TestesAutomatizados/Testify/erros"
)

func TestMedia(t *testing.T) {
	testes := []struct {
		input          []float64
		outputEsperado float64
		erroEsperado   error
	}{
		{
			input:          []float64{10, 20, 30},
			outputEsperado: 20,
			erroEsperado:   nil,
		},
		{
			input:          []float64{},
			outputEsperado: 0,
			erroEsperado:   erros.ErroListaVazia,
		},
		{
			input:          []float64{5, 5, 5, 5},
			outputEsperado: 5,
			erroEsperado:   nil,
		},
	}

	for _, teste := range testes {
		output, err := Media(teste.input)

		if err != teste.erroEsperado {
			t.Errorf("Erro esperado: %v, obtido: %v", teste.erroEsperado, err)
		}

		if output != teste.outputEsperado {
			t.Errorf("Media(%v) = %v; esperado %v", teste.input, output, teste.outputEsperado)
		}
	}
}

func TestMDC(t *testing.T) {

	testes := []struct {
		n1             int
		n2             int
		outputEsperado int
		erroEsperado   error
	}{
		{
			n1:             48,
			n2:             18,
			outputEsperado: 6,
			erroEsperado:   nil,
		},
		{
			n1:             101,
			n2:             103,
			outputEsperado: 1,
			erroEsperado:   nil,
		},
		{
			n1:             -10,
			n2:             5,
			outputEsperado: 0,
			erroEsperado:   erros.ErroNumeroNegativo,
		},
	}

	for _, teste := range testes {
		resultado, err := MDC(teste.n1, teste.n2)
		if err != teste.erroEsperado {
			t.Errorf("Erro esperado: %v, obtido: %v", teste.erroEsperado, err)
		}
		if resultado != teste.outputEsperado {
			t.Errorf("MDC(%d, %d) = %d; esperado %d", teste.n1, teste.n2, resultado, teste.outputEsperado)
		}
	}
}

func TestFibonacci(t *testing.T) {
	testes := []struct {
		numero         int
		outputEsperado []int
		erroEsperado   error
	}{
		{
			numero:         5,
			outputEsperado: []int{1, 1, 2, 3, 5},
			erroEsperado:   nil,
		},
		{
			numero:         0,
			outputEsperado: []int{},
			erroEsperado:   nil,
		},
		{
			numero:         10,
			outputEsperado: []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55},
			erroEsperado:   nil,
		},
		{
			numero:         -1,
			outputEsperado: nil,
			erroEsperado:   erros.ErroNumeroNegativo,
		},
	}

	for _, teste := range testes {
		result, err := Fibonacci(teste.numero)
		if err != teste.erroEsperado {
			t.Errorf("Erro esperado: %v, obtido: %v", teste.erroEsperado, err)
		}
		if !reflect.DeepEqual(result, teste.outputEsperado) {
			t.Errorf("Fibonacci(%d) = %v; esperado %v", teste.numero, result, teste.outputEsperado)
		}
	}
}
