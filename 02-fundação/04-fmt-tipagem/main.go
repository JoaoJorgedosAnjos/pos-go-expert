package main

import "fmt"

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
	fmt.Printf("O tipo de e é %T\n", e)
	fmt.Printf("O tipo de e é %v\n", e)
	fmt.Printf("O tipo de f é %T\n", f)
	fmt.Printf("O tipo de f é %v\n", f)

	/*
	   O tipo de e é float64
	   O tipo de e é 1.2
	   O tipo de f é main.ID
	   O tipo de f é 1
	*/

}

/*
### Visão Geral do Código Go

O código define e inicializa várias variáveis com tipos diferentes, incluindo um **tipo personalizado** (`ID`). O ponto principal do programa é usar o pacote `fmt` para **imprimir no console tanto o tipo quanto o valor** de duas dessas variáveis (`e` e `f`), mostrando como `fmt.Printf` pode ser usado para depuração e inspeção.

---

### Detalhes Importantes

* **`package main`**: Declara que este é um programa executável.
* **`import "fmt"`**: Importa o pacote `fmt`, essencial para formatação de entrada/saída (impressão no console).
* **`const a = "Hello, world"`**: Declara uma **constante** `a` do tipo `string`. Constantes são valores que não podem ser alterados após a compilação.
* **`type ID int`**: Cria um **tipo alias** chamado `ID` que, internamente, é um `int`. Isso ajuda a melhorar a legibilidade do código, indicando que variáveis do tipo `ID` são usadas como identificadores, mesmo que se comportem como inteiros.
* **Declaração de Variáveis (`var`)**:
    * `b bool = true`: Variável booleana.
    * `c int = 10`: Variável inteira.
    * `d string = "nome"`: Variável string.
    * `e float64 = 1.2`: Variável de ponto flutuante de 64 bits.
    * `f ID = 1`: Uma variável do **tipo `ID` personalizado**, inicializada com um valor inteiro.
* **`func main()`**: A função principal onde a execução do programa começa.
* **`fmt.Printf`**: Dentro de `main`, esta função é usada para:
    * Imprimir o **tipo** de `e` (`%T`) e seu **valor** (`%v`). A saída mostra `float64` e `1.2`.
    * Imprimir o **tipo** de `f` (`%T`) e seu **valor** (`%v`). A saída mostra `main.ID` (indicando que `ID` é um tipo definido no pacote `main`) e `1`.
    * O `\n` ao final de cada `Printf` garante que cada saída apareça em uma nova linha no terminal.

Em resumo, o código demonstra a declaração de constantes, a criação de tipos alias, a declaração e inicialização de variáveis de diferentes tipos básicos, e como usar `fmt.Printf` para inspecionar e exibir tanto o tipo quanto o valor de variáveis em Go.
*/
