package math

import "github.com/danubiobwm/cursoGolangAvac/cursoGolangAvac/TestesAutomatizados/Testify/erros"

func Media(numeros []float64) (float64, error) {
	if len(numeros) == 0 {
		return 0, erros.ErroListaVazia
	}

	soma := 0.0
	for _, numero := range numeros {
		soma += numero
	}

	return soma / float64(len(numeros)), nil
}

func MDC(n1, n2 int) (int, error) {
	if n1 < 0 || n2 < 0 {
		return 0, erros.ErroNumeroNegativo
	}

	for n2 != 0 {
		n1, n2 = n2, n1%n2
	}

	return n1, nil
}

func Fibonacci(n int) ([]int, error) {
	if n < 0 {
		return nil, erros.ErroNumeroNegativo
	}

	sequencia := make([]int, n)

	for i := 0; i < n; i++ {
		if i < 2 {
			sequencia[i] = 1
			continue
		}
		sequencia[i] = sequencia[i-1] + sequencia[i-2]
	}

	return sequencia, nil
}
