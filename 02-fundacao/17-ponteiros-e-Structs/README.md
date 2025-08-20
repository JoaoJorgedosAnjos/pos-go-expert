package main

import "fmt"

// Definição da struct Conta
type Conta struct {
    saldo int // Campo que armazena o saldo da conta
}

// Função construtora que retorna um ponteiro para uma nova Conta
func NewConta() *Conta {
    return &Conta{saldo: 0} // Retorna endereço de uma nova Conta com saldo 0
}

// Método que opera sobre uma Conta (receiver por ponteiro)
func (c *Conta) simular(valor int) int {
    c.saldo += valor        // Modifica o saldo original (por referência)
    println(c.saldo)        // Imprime o novo saldo
    return c.saldo          // Retorna o saldo atualizado
}

func main() {
    // Criação direta de uma Conta (por valor)
    conta := Conta{saldo: 100}
    
    // Chama o método simular
    conta.simular(200)
    
    // Imprime o saldo após a simulação
    println(conta.saldo)
}


Olá\! Vamos analisar esse código Go detalhadamente para entender o que está acontecendo, especialmente com o uso de **structs**, **ponteiros** e **métodos**.

-----

### A Struct `Conta`

```go
type Conta struct {
    saldo int
}
```

  * **O que é?** `Conta` é uma **`struct`**, uma estrutura de dados que agrupa campos relacionados. Neste caso, a `Conta` tem apenas um campo chamado `saldo`, que é do tipo `int` (número inteiro).
  * **Em essência:** Uma `struct` é como uma planta ou um molde para criar objetos. Por si só, ela não tem dados; ela apenas define a forma que os dados terão.

-----

### O Construtor `NewConta`

```go
func NewConta() *Conta{
    return &Conta{saldo:0}
}
```

  * **O que é?** Esta é uma função que atua como um **construtor**. Embora Go não tenha construtores no sentido tradicional de linguagens como Java ou C++, é uma convenção comum criar funções com o prefixo `New` para inicializar novas instâncias de uma `struct`.
  * **O que ele faz?** Ele retorna um **ponteiro** para uma nova `Conta`.
      * `Conta{saldo:0}` cria uma nova instância da `struct` `Conta` e inicializa o campo `saldo` com o valor `0`.
      * O operador **`&`** (endereço de memória) obtém o endereço de memória dessa nova `struct`.
  * **Por que um ponteiro?** Retornar um ponteiro (`*Conta`) é uma prática comum para evitar a cópia de grandes `structs` e para garantir que o código que recebe a `struct` possa modificar a original, e não apenas uma cópia.

-----

### O Método `simular`

```go
func (c *Conta) simular(valor int) int {
    c.saldo += valor
    println(c.saldo)
    return c.saldo
}
```

  * **O que é?** Esta é uma **função de método**. A sintaxe `(c *Conta)` indica que esta função está ligada à `struct` `Conta`. Mais especificamente, ela está ligada a um **ponteiro** para a `Conta` (`*Conta`).
  * **O que ele faz?**
    1.  `c.saldo += valor`: Ele modifica o campo `saldo` da `Conta` somando o `valor` passado.
    2.  `println(c.saldo)`: Ele imprime o novo saldo no console.
    3.  `return c.saldo`: Ele retorna o valor final do saldo.
  * **Por que o receptor é um ponteiro (`*Conta`)?**
      * O `c` é chamado de **receptor**.
      * Usar um ponteiro como receptor (`*Conta`) é crucial. Se o receptor fosse apenas uma cópia da `struct` (`c Conta`), a modificação (`c.saldo += valor`) só afetaria a cópia dentro do método, e o `saldo` da `Conta` original em `main` permaneceria inalterado.
      * Com um ponteiro, o método `simular` acessa e modifica diretamente a `Conta` original na memória, garantindo que a alteração seja permanente.

-----

### A Função Principal `main`

```go
func main() {
    conta := Conta{saldo:100}
    conta.simular(200)
    println(conta.saldo)
}
```

  * `conta := Conta{saldo:100}`: Uma nova `struct` `Conta` é criada e inicializada com um `saldo` de `100`. Note que aqui você está criando uma **variável de valor**, ou seja, a variável `conta` armazena a `struct` diretamente, e não um ponteiro para ela.
  * `conta.simular(200)`: O método `simular` é chamado. O Go é inteligente e, mesmo que a variável `conta` seja uma `struct` de valor, ele automaticamente a converte para um ponteiro (`&conta`) para chamar o método, pois o método `simular` espera um ponteiro como receptor.
  * `println(conta.saldo)`: Acontece a mágica. Como o método `simular` modificou o `saldo` da `Conta` original, o valor impresso aqui será o saldo final, que é `300`.

### A Saída do Programa

O resultado da execução deste código será:

```
300
300
```

1.  O primeiro `300` vem do `println` dentro do método `simular`.
2.  O segundo `300` vem do `println` na função `main`, provando que o `saldo` da `conta` original foi alterado.