
# Entendendo Funções Variádicas em Go: Um Exemplo Prático

Este capítulo explora um conceito fundamental em Go: funções variádicas. Através de um exemplo simples, demonstraremos como criar e utilizar funções que podem aceitar um número variável de argumentos.

## 1. Estrutura Básica de um Programa Go

Todo programa Go começa com a declaração do pacote `main` e a função `main`.

```go
package main

import (
    "fmt"
)

func main() {
    // Nosso código principal reside aqui
}
```

* **`package main`**: Declara que este arquivo pertence ao pacote `main`. Em Go, programas executáveis devem sempre estar no pacote `main`.
* **`import ("fmt")`**: Importa o pacote `fmt` (format). Este pacote fornece funções para formatação de entrada e saída, como a impressão de texto no console.
* **`func main() { ... }`**: Esta é a função principal onde a execução do programa começa. Quando você executa um programa Go, a primeira função a ser chamada é `main`.

## 2. A Função `sum`: Exemplo de Função Variádica

A peça central do nosso exemplo é a função `sum`. Ela é projetada para calcular a soma de qualquer quantidade de números inteiros.

```go
func sum(numeros ...int) int {
    total := 0
    for _, numero := range numeros {
        total += numero
    }
    return total
}
```

Vamos analisar cada parte:

* **`func sum(...)`**: Declara uma nova função chamada `sum`.
* **`numeros ...int`**: Esta é a parte crucial que torna `sum` uma **função variádica**.
    * Os três pontos (`...`) antes do tipo (`int`) indicam que `sum` pode aceitar zero ou mais argumentos do tipo `int`.
    * Internamente, esses argumentos são coletados em um *slice* de inteiros chamado `numeros`. Um slice é uma fatia dinâmica de um array, permitindo armazenar uma coleção de elementos do mesmo tipo.
* **`int` (após os parênteses)**: Indica que a função `sum` retornará um valor do tipo `int`.

### Lógica Interna da Função `sum`

```go
    total := 0
    for _, numero := range numeros {
        total += numero
    }
    return total
```

1.  **`total := 0`**: Uma variável chamada `total` é inicializada com o valor `0`. Esta variável armazenará a soma acumulada dos números.
2.  **`for _, numero := range numeros { ... }`**: Este é um laço `for...range` comum em Go, usado para iterar sobre coleções.
    * `range numeros`: Itera sobre cada elemento no slice `numeros`.
    * `_`: O primeiro valor retornado por `range` é o índice do elemento. Usamos o *identificador em branco* (`_`) para indicar que não estamos interessados no índice neste caso.
    * `numero`: O segundo valor retornado é o próprio elemento do slice. Em cada iteração, `numero` conterá um dos inteiros passados para a função `sum`.
3.  **`total += numero`**: Em cada iteração, o valor atual de `numero` é adicionado ao `total`.
4.  **`return total`**: Após percorrer todos os números no slice, a função retorna o valor final de `total`, que é a soma de todos os argumentos.

## 3. Chamando a Função `sum` na `main`

De volta à nossa função `main`, é onde chamamos `sum` e exibimos o resultado:

```go
func main() {
    fmt.Println(sum(1, 34, 43, 54))
}
```

* **`sum(1, 34, 43, 54)`**: Aqui, a função `sum` é chamada com quatro argumentos inteiros. Graças à sua natureza variádica, ela aceita esses números sem problemas.
* **`fmt.Println(...)`**: A função `Println` do pacote `fmt` é usada para imprimir o valor retornado por `sum` no console, seguido por uma nova linha.

## 4. Executando o Código

Quando você executa este programa:

1.  A função `main` é iniciada.
2.  `sum(1, 34, 43, 54)` é chamada.
3.  Dentro de `sum`, o slice `numeros` será `[1, 34, 43, 54]`.
4.  O laço `for` somará os elementos: `0 + 1 = 1`, `1 + 34 = 35`, `35 + 43 = 78`, `78 + 54 = 132`.
5.  A função `sum` retorna `132`.
6.  `fmt.Println(132)` imprime `132` na saída padrão.

## Conclusão

Este exemplo demonstra a poderosa funcionalidade das funções variádicas em Go. Elas permitem escrever código mais flexível e reutilizável, capaz de lidar com um número variável de entradas. A combinação de funções variádicas com loops `for...range` é um padrão comum para processar coleções de dados de forma concisa e eficiente em Go.

---