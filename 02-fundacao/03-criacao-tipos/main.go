package main

const a = "Hello, world"

type ID int // Declaração de tipo alias

var (
	b bool    = true
	c int     = 10
	d string  = "nome"
	e float64 = 1.2
	f ID      = 1
)

func main() {
	//a := "X" //var a string
	//a = "Z"  //sem : quando for reatribuir, := só na primeira vez
	println(f)
}

/*
fmt significa "format" (formato)

Em Go, o pacote fmt é uma abreviação de "format" ou "formatted I/O" (entrada/saída formatada).

Ele é o pacote padrão da biblioteca Go que oferece funcionalidades para:

    Formatar dados (números, strings, booleanos, etc.) para exibição ou para construir strings.
    Imprimir esses dados na saída padrão (geralmente o terminal).
    Ler dados da entrada padrão (geralmente o teclado).

A ideia principal é que você pode controlar exatamente como seus dados aparecem quando são impressos ou transformados em uma string, daí o "formato". Funções como fmt.Printf() são um ótimo exemplo disso, onde você usa verbos de formatação (como %s, %d, %f) para definir o layout da saída.
*/