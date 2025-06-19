package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Pessoa interface {
	//interface no Go só permite passar assinatura de métodos
	Desativar()
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado", c.Nome)
}

func Desativacao(pessoa Pessoa){
	pessoa.Desativar()
}

func main() {
	joao := Cliente{
		Nome:  "João",
		Idade: 21,
		Ativo: true,
	}

	joao.Cidade = "Brasília"
	//joao.Desativar()
	Desativacao(joao)
}
