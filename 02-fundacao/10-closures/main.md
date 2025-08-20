# Explorando Funções Anônimas e Closures em Go

Este capítulo baseia-se no nosso entendimento de funções variádicas, introduzindo um conceito poderoso em Go: **funções anônimas** (ou funções literais) e como elas podem ser usadas em conjunto com outras funções.

## 1. Estrutura Base do Programa Go

Como já sabemos, todo programa Go começa com o pacote `main` e a função `main` é o ponto de entrada da execução.

```go
package main

import (
    "fmt"
)

func main() {
    // Nosso código principal reside aqui
}
```

* **`package main`**: Indica que este é um programa executável.
* **`import ("fmt")`**: Importa o pacote `fmt` para operações de entrada e saída.

## 2. A Função Variádica `sum`

A função `sum` permanece a mesma do nosso exemplo anterior. Ela aceita um número variável de argumentos inteiros e retorna a soma deles.

```go
func sum(numeros ...int) int {
    total := 0
    for _, numero := range numeros {
        total += numero
    }
    return total
}
```

* **`func sum(numeros ...int) int`**: Declara a função `sum` que recebe um **slice variádico** de inteiros (`...int`) e retorna um `int`.
* A lógica interna **itera** sobre o slice `numeros` e acumula a soma em `total`, que é então retornado.

## 3. O Destaque: Funções Anônimas na `main`

A principal mudança neste código está na função `main`, onde introduzimos uma **função anônima**.

```go
func main() {
    total := func () int {
        return sum(1, 34, 43, 54) * 2
    }()
    fmt.Println(total)
}
```

Vamos desmembrar essa parte:

* **`func () int { ... }`**: Esta é a **declaração de uma função anônima**.
    * **`func`**: Palavra-chave para declarar uma função.
    * **`()`**: Indica que esta função não recebe nenhum argumento.
    * **`int`**: Indica que esta função anônima retornará um valor inteiro.
    * **`{ return sum(1, 34, 43, 54) * 2 }`**: Este é o **corpo** da função anônima.
        * Dentro dela, chamamos nossa função `sum` com os números `1, 34, 43, 54`.
        * O resultado de `sum` é então **multiplicado por 2**.
* **`()` (após o bloco da função anônima)**: Este par de parênteses é crucial! Ele **executa imediatamente** a função anônima que acabamos de definir. Sem esses parênteses, estaríamos apenas definindo a função, não chamando-a.
* **`total := ...`**: O valor retornado pela execução imediata da função anônima é atribuído à variável `total`.
* **`fmt.Println(total)`**: Finalmente, o valor de `total` é impresso no console.

### Por que usar uma Função Anônima aqui?

Neste exemplo específico, a função anônima pode parecer um pouco redundante, pois poderíamos ter escrito `total := sum(1, 34, 43, 54) * 2` diretamente. No entanto, funções anônimas são extremamente úteis em cenários mais complexos, como:

* **Closures**: Funções anônimas podem "capturar" e acessar variáveis do escopo em que foram criadas, mesmo após o escopo original ter terminado. Isso é conhecido como **closure**.
* **Callbacks**: Passar uma função como argumento para outra função.
* **Goroutines**: Iniciar novas threads de execução de forma concisa.
* **Inicialização única**: Como visto aqui, para encapsular uma operação que deve ser executada apenas uma vez no momento da declaração.

## 4. Fluxo de Execução

Ao executar este programa:

1.  A função `main` é iniciada.
2.  Uma função anônima é definida.
3.  Imediatamente após sua definição, essa função anônima é executada:
    * Ela chama `sum(1, 34, 43, 54)`, que retorna `132`.
    * O resultado `132` é multiplicado por `2`, resultando em `264`.
    * O valor `264` é retornado pela função anônima.
4.  O valor `264` é atribuído à variável `total`.
5.  `fmt.Println(total)` imprime `264` no console.

## Conclusão

Este exemplo demonstra como as **funções anônimas** podem ser definidas e executadas em Go. Elas oferecem flexibilidade para escrever código mais conciso e expressivo, especialmente quando precisamos de uma função para uma tarefa específica e de uso único, ou quando queremos criar closures. Compreender as funções anônimas é um passo importante para dominar padrões de programação mais avançados em Go.

---