package main

import "fmt"

func main() {
	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	for i := 0; i < len(meuArray); i++ {
		fmt.Printf("%v\n", meuArray[i])
	}

	for i, v := range meuArray {
		fmt.Printf("O valor do indice %v é %v\n", i, v)
	}

	for i, v := range meuArray {
		fmt.Printf("O valor do indice %d é %d\n", i, v)
	}
	/*
	for i (index[0] = 0), v (value[0] = 10)range (por todo o tamanho do) meuArray
	*/

	//fmt.Printf("%v", meuArray[1])
	//fmt.Println(len(meuArray) -1) // tamanho do array começando do 1 print = 20

	//fmt.Println(meuArray[len(meuArray) -1]) //retonar o ultimo valor do array
}

/*
package main

import "fmt"

func main() {
    numero := 123
    texto := "abc"

    fmt.Printf("Usando %%d para numero: %d\n", numero) // OK, numero é int
    fmt.Printf("Usando %%v para numero: %v\n", numero) // OK, %v imprime o int

    // fmt.Printf("Usando %%d para texto: %d\n", texto) // ERRO DE COMPILAÇÃO! %d espera um int
    fmt.Printf("Usando %%v para texto: %v\n", texto)   // OK, %v imprime o string
}

Qual a melhor forma de fazer um for?

Em Go, a "melhor" forma de fazer um for depende do que você precisa fazer. Existem três formas principais, e cada uma tem seu caso de uso ideal:

    for com inicializador, condição e pós-declaração (o for tradicional):
    Go

for i := 0; i < len(meuArray); i++ {
    fmt.Printf("%v\n", meuArray[i])
}

    Quando usar:
        Quando você precisa de um contador (i) que é usado diretamente nos cálculos ou para acessar elementos por índice.
        Quando você precisa de controle preciso sobre a iteração (ex: pular alguns índices, iterar para trás).
        Quando você está trabalhando com arrays onde o índice é tão importante quanto o valor, ou mais.

for...range (para iterar sobre coleções):
Go

for i, v := range meuArray { // i é o índice, v é o valor
    fmt.Printf("O valor do índice %v é %v\n", i, v)
}

    Quando usar:
        Esta é a forma mais idiomática e geralmente preferida para iterar sobre slices, arrays, strings, maps e channels.
        Quando você precisa tanto do índice quanto do valor de cada elemento.
        Quando a legibilidade é uma prioridade, pois é mais conciso.
        Se você só precisa do valor, pode descartar o índice usando _: for _, v := range meuArray.
        Se você só precisa do índice, pode descartar o valor: for i := range meuArray.

for como while (loop condicional):
Go

    i := 0
    for i < 5 { // Loop continuará enquanto i for menor que 5
        fmt.Println(i)
        i++
    }

        Quando usar:
            Quando você precisa de um loop que continua enquanto uma condição for verdadeira, sem um contador explícito ou uma coleção para iterar.
            Para implementar loops "infinitos" (for {}).

Conclusão sobre a "melhor forma" do for:

    Para iterar sobre a maioria das coleções (slices, arrays, maps, strings), o for...range é quase sempre a melhor e mais idiomática escolha em Go. Ele é mais legível, menos propenso a erros de "off-by-one" (erro de um) e mais conciso.

    O for tradicional com i := 0; i < len; i++ ainda é útil para cenários onde você precisa do controle exato do índice para operações matemáticas complexas, ou quando você está construindo algo que não é uma iteração simples sobre uma coleção.

    O for como while é para loops baseados puramente em uma condição booleana.


*/
