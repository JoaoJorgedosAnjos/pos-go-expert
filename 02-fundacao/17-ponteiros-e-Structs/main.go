package main

import "fmt"

type Cliente struct {
	nome string
}

func (c Cliente) andou() {
	c.nome = "Wesley Willinas"
	fmt.Printf("O cliente %v andou\n", c.nome)
}

func main() {
	wesley := Cliente{
		nome: "Wesley",
	}

	wesley.andou()
	fmt.Printf("O valor da struct com nome %v\n", wesley.nome)
}
