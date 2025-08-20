# Explorando `Structs`, Composição e Métodos em Go

Este capítulo aborda a maneira como Go organiza dados usando `structs`, a poderosa técnica de **composição** para construir tipos mais complexos, e a distinção crucial entre **receivers por valor** e **receivers por ponteiro** em métodos.

## 1. Estrutura Básica de um Programa Go

Todo programa Go começa com o pacote `main` e a função `main()`.

```go
package main

import "fmt" // Importa o pacote fmt para funções de formatação (como imprimir no console)

func main() {
    // A execução do seu programa Go inicia aqui
}
```

---

## 2. Definindo `Structs`: `Endereco` e `Cliente`

Em Go, uma `struct` (estrutura) é um tipo de dado composto que agrupa campos (variáveis) de diferentes tipos sob um único nome. Ela é a maneira de representar entidades complexas do mundo real.

### 2.1. A `struct Endereco`

```go
type Endereco struct {
    Logradouro string
    Numero     int
    Cidade     string
    Estado     string
}
```

* **`type Endereco struct { ... }`**: Declara um novo tipo chamado `Endereco`.
* Esta `struct` contém quatro campos para armazenar informações de um endereço: `Logradouro` (rua, avenida, etc.), `Numero`, `Cidade` e `Estado`. Todos os campos são do tipo `string`, exceto `Numero`, que é `int`.

### 2.2. A `struct Cliente` com Composição

```go
type Cliente struct {
    Nome  string
    Idade int
    Ativo bool
    Endereco // Composição por embedding (incorporação anônima)
}
```

A `struct Cliente` é mais interessante, pois demonstra a **composição**.

* **`Nome string`, `Idade int`, `Ativo bool`**: São campos diretos da `struct Cliente`.
* **`Endereco`**: Esta linha é um exemplo de **embedding** (incorporação anônima). Em vez de `Cliente` ter um campo como `MeuEndereco Endereco`, simplesmente listamos o tipo `Endereco`.
    * **Benefício**: Os campos da `struct Endereco` (como `Cidade`, `Estado`) são "promovidos" e podem ser acessados diretamente a partir de uma instância de `Cliente`. Por exemplo, `joao.Cidade` é um atalho para `joao.Endereco.Cidade`. Go faz essa mágica para você!
    * Esta é a forma de Go de alcançar a **reutilização de código** e construir relacionamentos "tem um" (Cliente *tem um* Endereço) sem a complexidade da herança tradicional.

---

## 3. O Método `Desativar()`: Entendendo os Receivers

Em Go, você pode anexar **métodos** a qualquer tipo definido pelo usuário, incluindo `structs`. Um método é simplesmente uma função com um argumento especial chamado **receiver** (receptor).

```go
func (c Cliente) Desativar() { // <- O receiver "c Cliente" é por VALOR
    c.Ativo = false
    fmt.Printf("O cliente %s foi desativado\n", c.Nome)
}
```

* **`func (c Cliente) Desativar()`**: Esta linha define um método chamado `Desativar` para o tipo `Cliente`.
    * **`c Cliente`**: Este é o **receiver por valor**. Isso significa que, quando você chama `joao.Desativar()`, o Go cria uma **cópia completa** da `struct joao` e passa essa cópia para o método.
* **`c.Ativo = false`**: Esta linha tenta modificar o campo `Ativo` do cliente.
    * **O Problema**: Como `c` é uma **cópia**, a modificação de `c.Ativo` só afeta a cópia dentro do método. A `struct Cliente` original (`joao` em `main`) permanece **inalterada**.
* **`fmt.Printf(...)`**: Esta linha imprime uma mensagem confirmando a "desativação", mas lembre-se que isso se refere à cópia dentro do método.

---

## 4. A Função `main()`: Orquestrando o Programa

A função `main()` é onde seu programa começa e onde as `structs` são instanciadas e os métodos são chamados.

```go
func main() {
    joao := Cliente{
        Nome:  "João",
        Idade: 21,
        Ativo: true, // João é inicializado como ativo
    }

    // Acessando e modificando campos da struct Cliente
    joao.Cidade = "Brasília" // Funciona devido à promoção do campo 'Cidade' da struct Endereco
                             // É o mesmo que joao.Endereco.Cidade = "Brasília"

    // Chamando o método Desativar
    joao.Desativar() // Esta chamada executa o método com a CÓPIA de 'joao'

    // Verificando o estado de 'joao' APÓS a chamada do método
    fmt.Printf("Status final do cliente %s: Ativo = %t\n", joao.Nome, joao.Ativo)
}
```

### O que Acontece na Execução:

1.  Uma instância de `Cliente` chamada `joao` é criada, e seu campo `Ativo` é definido como `true`.
2.  `joao.Cidade = "Brasília"`: O campo `Cidade` (promovido de `Endereco`) da `joao` original é alterado para "Brasília". **Isso funciona.**
3.  `joao.Desativar()`:
    * Uma **cópia de `joao`** é feita e passada para o método `Desativar()`.
    * Dentro de `Desativar()`, o campo `Ativo` da **cópia** é alterado para `false`.
    * A mensagem "O cliente João foi desativado" é impressa (usando os dados da cópia).
    * O método termina, e a **cópia é descartada**. A `joao` original na `main` **permanece com `Ativo` igual a `true`**.
4.  `fmt.Printf("Status final...")`: Quando esta linha é executada, ela imprime o estado do `joao` **original**, que, como vimos, ainda tem `Ativo` como `true`.

---

## Conclusão: Receivers por Valor vs. Receivers por Ponteiro

Este código é um excelente exemplo para entender a diferença crucial entre **receivers por valor** e **receivers por ponteiro** em Go:

* **Receiver por Valor (`(c Cliente)`):** O método opera em uma **cópia** da `struct`. Quaisquer modificações dentro do método não afetam a `struct` original.
* **Receiver por Ponteiro (`(c *Cliente)`):** Se você deseja que um método **modifique** a `struct` original (em vez de uma cópia), você deve usar um ponteiro como receiver. Isso passa uma referência à localização da `struct` na memória, permitindo que o método a altere diretamente.

Para que seu método `Desativar()` realmente altere o estado do cliente, você precisaria modificá-lo para:

```go
func (c *Cliente) Desativar() { // Note o asterisco (*) aqui
    c.Ativo = false
    fmt.Printf("O cliente %s foi desativado\n", c.Nome)
}
```

Ao compreender essa distinção, você pode escrever código Go mais eficaz e prever seu comportamento corretamente!