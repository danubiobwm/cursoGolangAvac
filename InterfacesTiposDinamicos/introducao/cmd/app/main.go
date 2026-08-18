package main

import (
	conversor "github.com/danubiobwm/cursoGolangAvac/cursoGolangAvac/InterfacesTiposDinamicos/introducao/pkg/converson"
	"github.com/danubiobwm/cursoGolangAvac/cursoGolangAvac/InterfacesTiposDinamicos/introducao/pkg/converson/json"
	"github.com/danubiobwm/cursoGolangAvac/cursoGolangAvac/InterfacesTiposDinamicos/introducao/pkg/converson/xml"
)

func main() {
	gerenciador := conversor.NovoGerenciador()
	gerenciador.RegistrarConversor(json.ConversorJSON{})
	gerenciador.RegistrarConversor(xml.ConversorXML{})

	dadosJSON := `{
		"pessoa": {
			"nome": "João",
			"idade": 30
		}
	}`

	dadosXML := `
	<pessoa>
		<nome>João</nome>
		<idade>30</idade>
	</pessoa>`

	gerenciador.Processar("JSON", dadosJSON)
	gerenciador.Processar("XML", dadosXML)
}
