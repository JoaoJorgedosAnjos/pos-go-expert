package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome   string
	Idade  int
	Ativo  bool
	Endereco
	//ou
	//Adress Endereco
	/*
		type Endereco struct { // size=56 (0x38)
	    Logradouro string
	    Numero     int
	    Cidade     string
	    Estado     string
		main.Endereco on pkg.go.dev
	}
	*/
}

func main() {
	joao := Cliente{
		Nome:  "João",
		Idade: 21,
		Ativo: true,
	}

	joao.Ativo = false
	//joao.Adress.Cidade = "Brasília"
	joao.Cidade = "Brasília"
	//ou
	//joao.Endereco.Cidade = "Brasília"
	fmt.Println(joao.Nome)
}
