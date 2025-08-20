package main // Declara o pacote principal do programa. Em Go, programas executáveis devem estar no pacote 'main'.

import "fmt" // Importa o pacote "fmt", que fornece funções para formatação de entrada e saída (como imprimir no console).

func main() { // Declara a função principal (main), que é o ponto de entrada de qualquer programa Go executável.
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100} // Declara e inicializa um slice de inteiros chamado 's' com os valores 2, 4, 6, 8 e 10.
	// Um slice é uma visão flexível de um array.

	// Imprime o comprimento (len), capacidade (cap) e os elementos do slice 's' original.
	// len() retorna o número de elementos no slice.
	// cap() retorna a capacidade do array subjacente (o número máximo de elementos que o slice pode conter sem realocação).
	// %d é um verbo de formatação para inteiros decimais.
	// %v é um verbo de formatação para o valor em seu formato padrão.
	fmt.Printf("len=%d cap=%d v=%v\n", len(s), cap(s), s)

	// Imprime o comprimento, capacidade e os elementos de um novo slice criado a partir de 's'.
	// s[:0] cria um slice com comprimento 0, mas com a mesma capacidade do slice original 's'.
	// Ele aponta para o início do array subjacente de 's'.
	fmt.Printf("len=%d cap=%d v=%v\n", len(s[:0]), cap(s[:0]), s[:0])

	// Imprime o comprimento, capacidade e os elementos de um novo slice criado a partir de 's'.
	// s[:4] cria um slice que inclui elementos do índice 0 até o índice 3 (4 elementos no total).
	// Sua capacidade será a mesma do slice original 's'.
	fmt.Printf("len=%d cap=%d v=%v\n", len(s[:4]), cap(s[:4]), s[:4])

	// Imprime o comprimento, capacidade e os elementos de um novo slice criado a partir de 's'.
	// s[2:] cria um slice que começa a partir do índice 2 (o terceiro elemento) até o final do slice original 's'.
	// Sua capacidade será a capacidade restante do array subjacente a partir do índice 2.
	fmt.Printf("len=%d cap=%d v=%v\n", len(s[2:]), cap(s[2:]), s[2:])

	s = append(s, 110) // Adiciona o valor 12 ao final do slice 's'.
	// Se a capacidade do slice for excedida, Go realocará um novo array subjacente maior.

	// Imprime o comprimento, capacidade e os elementos de um novo slice criado a partir do 's' modificado.
	// Este slice também começa a partir do índice 2 até o final.
	fmt.Printf("len=%d cap=%d v=%v\n", len(s[2:]), cap(s[2:]), s[2:])
}
