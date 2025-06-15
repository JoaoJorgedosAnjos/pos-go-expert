package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	joao := Cliente{
		Nome:  "João",
		Idade: 21,
		Ativo: true,
	}

	joao.Ativo = false
	fmt.Println(joao.Nome)
}
