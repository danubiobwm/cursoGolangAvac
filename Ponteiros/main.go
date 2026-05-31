package main

func main() {
	var a int = 10
	var b *int = &a
	println("Valor de a:", a)
	println("Endereço de a:", &a)
	println("Valor de b (endereço de a):", b)
	println("Valor apontado por b (valor de a):", *b)

	pessoa := NovaPessoa("João", 30)
	println("Nome:", pessoa.Nome)
	println("Idade:", pessoa.Idade)
	println("Telefone:", pessoa.Telefone())

	pessoa.AtualizarIdade(31)
	pessoa.AtualizarTelefone("123-456-7890")
	println("Nome:", pessoa.Nome)
	println("Idade atualizada:", pessoa.Idade)
	println("Telefone atualizado:", pessoa.Telefone())
}

type Pessoa struct {
	Nome     string
	Idade    int
	telefone *string
}

func NovaPessoa(nome string, idade int) Pessoa {
	return Pessoa{
		Nome:  nome,
		Idade: idade,
	}
}

func (p Pessoa) Telefone() string {
	if p.telefone == nil {
		return ""
	}
	return *p.telefone
}

func (p *Pessoa) AtualizarIdade(novaIdade int) {
	p.Idade = novaIdade
}

func (p *Pessoa) AtualizarTelefone(novoTelefone string) {
	p.telefone = &novoTelefone
}
