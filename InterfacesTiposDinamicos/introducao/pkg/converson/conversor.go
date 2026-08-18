package conversor

import "fmt"

type (
	Conversor interface {
		Converter(string) (map[string]any, error)
		Formato() string
	}

	Gerenciador struct {
		conversores map[string]Conversor
	}
)

func NovoGerenciador() *Gerenciador {
	return &Gerenciador{
		conversores: make(map[string]Conversor),
	}
}

func (g *Gerenciador) RegistrarConversor(c Conversor) {
	g.conversores[c.Formato()] = c
}

func (g *Gerenciador) Processar(formato, dados string) {
	conversor, ok := g.conversores[formato]
	if !ok {
		fmt.Printf("Nenhum conversor cadastrado com o formato %s\n", formato)
		return
	}

	resultado, err := conversor.Converter(dados)
	if err != nil {
		fmt.Printf("Erro convertendo formato %s\n", formato)
		return
	}

	fmt.Printf("Dados processados (%s) - %v\n", formato, resultado)
}
