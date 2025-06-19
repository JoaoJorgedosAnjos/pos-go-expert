package main

func main() {
	// Entendendo a Memória e as Variáveis em Go

	/*
	   Ao declarar uma variável em Go (ou em qualquer linguagem de programação de baixo nível, como C/C++), o que acontece nos bastidores é fascinante:

	   1.  **Reserva de Espaço na Memória**: O sistema operacional, a pedido do programa, **reserva um pequeno "pedaço" ou "espaço" na memória RAM** do computador para armazenar o valor da sua variável. Esse espaço tem um **endereço** único (como um CEP) onde ele está localizado.

	   2.  **Atribuição de Valor**: O valor que você atribui à variável (neste caso, o número 10) é então **gravado nesse espaço de memória** que foi reservado.

	   3.  **Acesso ao Valor**: Sempre que você precisa "usar" o valor dessa variável no seu código, o programa **acessa o endereço de memória** onde ela está guardada, "lê" o valor que está lá dentro e o utiliza na operação desejada.

	   Em resumo: a variável (ex: 'a') é um **nome amigável** que usamos no código para nos referirmos a um **endereço específico na memória**, onde o **valor** real daquela variável está armazenado.
	*/

	a := 10 // Declaramos a variável 'a' e atribuímos o valor 10 a ela.
	// Isso aciona o processo de reserva de memória e armazenamento do valor.
	//println(&a) printa o endereço da memória, não o valor da variável
	var ponteiro *int = &a
	println(ponteiro)  //acessa o endereço da memória
	println(*ponteiro) //acessa o valor do endereço
	/* ou
	ponteiro := &a
	*/
	*ponteiro = 20
	println(a)
	b := &a
	*b = 30
	println(b)
	println(*b)
}
