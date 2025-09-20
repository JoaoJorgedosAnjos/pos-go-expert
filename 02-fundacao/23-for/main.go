package main

import "fmt"

func main() {
	/*
		1. O Loop for tradicional

		É o mais comum, similar ao for em C ou Java, com três componentes: uma instrução de inicialização, uma condição e uma instrução de pós-execução.
	*/
	// A instrução de inicialização é executada uma vez no início.
	// A condição é verificada a cada iteração.
	// A pós-instrução é executada após cada iteração.
	fmt.Printf("Loop 1\n")
	for i := 0; i < 5; i++ {
		fmt.Println("Número:", i)
	}

	/*
				2. O Loop for como while

		Nesta forma, usamos apenas a condição. O loop continua rodando enquanto a condição for true. A variável de controle precisa ser declarada e atualizada fora do for.
	*/
	// Declara a variável fora do loop.
	j := 0
	// O loop continua enquanto i for menor que 5.
	fmt.Printf("\nLoop 2\n")
	for j < 5 {
		fmt.Println("Número:", j)
		// Atualiza a variável dentro do loop.
		j++
	}
	fmt.Printf("\nLoop 3\n")

	/*
				3. O Loop for infinito

		Este loop não tem inicialização, condição ou pós-instrução. Ele roda para sempre, a menos que você use um comando como break para sair dele.
	*/
	contador := 0
	// Este loop é infinito, mas usaremos 'break' para sair.
	for {
		fmt.Println("Contador:", contador)
		contador++
		// A condição 'if' com 'break' é o que impede o loop de ser infinito de fato.
		if contador > 3 {
			break // Sai do loop quando o contador for maior que 3.
		}
	}

	/*
						4. O Loop for...range

		Usamos o for...range para iterar sobre elementos de coleções como slices, arrays, maps e strings. Ele retorna dois valores a cada iteração.
	*/
	// Loop sobre um slice
	frutas := []string{"maçã", "banana", "laranja"}
	// 'i' recebe o índice e 'fruta' recebe o valor do elemento.
	for i, fruta := range frutas {
		fmt.Println("Índice:", i, "Fruta:", fruta)
	}

	// Loop sobre um map
	estudantes := map[string]int{"João": 25, "Maria": 30}
	// 'nome' recebe a chave e 'idade' recebe o valor.
	for nome, idade := range estudantes {
		fmt.Println("Nome:", nome, "Idade:", idade)
	}

	numeros := []string{"um", "dois", "três"}
	for k, v := range numeros {
		println(k, v)
	}
	/*
	0 um
	1 dois
	2 três
	
	numeros := []string{"um", "dois", "três"}
	for k := range numeros {
		println(k)
	}
	apenas os index
	
	0 
	1 
	2 

	numeros := []string{"um", "dois", "três"}
	for _ , v := range numeros {
		println(v)
	}
	apenas os values
	um
	dois
	três
	*/
}
