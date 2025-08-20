## Entendendo o Código Go: Clientes, Endereços e Interfaces

Este código Go simples demonstra conceitos fundamentais como **structs**, **métodos** e **interfaces**. Vamos ver cada parte.

### Structs: Moldes para Seus Dados

No Go, **structs** são como "moldes" ou "plantas" para criar seus próprios tipos de dados complexos, agrupando várias informações sob um único nome.

1.  **`Endereco` (Endereço)**:
    ```go
    type Endereco struct {
        Logradouro string
        Numero     int
        Cidade     string
        Estado     string
    }
    ```
    Aqui, criamos uma struct `Endereco` para representar um endereço, com campos para logradouro, número, cidade e estado. Cada campo tem um tipo específico (string ou int).

2.  **`Cliente`**:
    ```go
    type Cliente struct {
        Nome  string
        Idade int
        Ativo bool
        Endereco // Struct incorporada (embedded)
    }
    ```
    A struct `Cliente` representa um cliente. Ela tem campos para `Nome`, `Idade` e um `Ativo` (booleano que indica se o cliente está ativo ou não). O mais interessante aqui é a linha `Endereco`. Isso é uma **struct incorporada (embedded struct)**. Significa que um `Cliente` automaticamente "ganha" todos os campos da struct `Endereco` sem que precisemos listá-los um por um. Por exemplo, em `main`, você pode acessar `joao.Cidade` diretamente, mesmo que `Cidade` seja um campo de `Endereco`.

### Interfaces: Definindo Comportamentos

Uma **interface** em Go define um "contrato": ela especifica quais métodos um tipo deve ter para "implementar" essa interface. Interfaces se preocupam com **o que** um tipo pode fazer, não com **como** ele é feito.

```go
type Pessoa interface {
    Desativar() // Qualquer tipo que tiver um método Desativar() implementa Pessoa
}
```
Aqui, a interface `Pessoa` declara que qualquer tipo que se comporte como uma "Pessoa" deve ter um método chamado `Desativar()` que não recebe argumentos nem retorna valores. Notou o comentário no código? `//interface no Go só permite passar assinatura de métodos`. Isso é crucial: interfaces listam apenas as assinaturas dos métodos, não a sua implementação.

### Métodos: Funções Associadas a Tipos

Um **método** é uma função especial que é "anexada" a um tipo específico (uma struct, por exemplo).

```go
func (c Cliente) Desativar() {
    c.Ativo = false
    fmt.Printf("O cliente %s foi desativado", c.Nome)
}
```
Esta é a implementação do método `Desativar()` para a struct `Cliente`.
* `func (c Cliente)`: O `(c Cliente)` é o **receptor** do método. Ele indica que este método `Desativar()` pertence ao tipo `Cliente`. Quando você chama `joao.Desativar()`, o `joao` é passado como o receptor `c` para dentro da função.
* `c.Ativo = false`: Dentro do método, acessamos o campo `Ativo` da instância de `Cliente` (que é uma **cópia** do objeto original) e o definimos como `false`.
* `fmt.Printf(...)`: Imprime uma mensagem no console, usando o nome do cliente.

**Ponto Importante**: Com o receptor `(c Cliente)` (chamado de **receptor de valor**), o método trabalha com uma **cópia** do objeto `Cliente`. Isso significa que, embora a mensagem seja impressa, a propriedade `Ativo` do objeto `joão` *original* na função `main` **não será alterada** para `false`. Se quiséssemos que a mudança fosse permanente no objeto original, o receptor deveria ser um **receptor de ponteiro**: `func (c *Cliente) Desativar()`.

### Função `Desativacao`: Usando a Interface

```go
func Desativacao(pessoa Pessoa){
    pessoa.Desativar()
}
```
Esta função é um exemplo claro do poder das interfaces. Ela aceita qualquer coisa que implemente a interface `Pessoa` (ou seja, qualquer tipo que tenha um método `Desativar()`). Dentro da função, ela simplesmente chama o método `Desativar()` no argumento `pessoa`. Ela não precisa saber se `pessoa` é um `Cliente` ou qualquer outro tipo, apenas que ele pode ser "desativado". Isso promove código mais flexível e reutilizável.

### `main`: Onde Tudo Acontece

```go
func main() {
    joao := Cliente{
        Nome:  "João",
        Idade: 21,
        Ativo: true,
    }

    joao.Cidade = "Brasília" // Acessando campo da struct Endereco incorporada
    Desativacao(joao)       // Passamos joao (que é um Cliente e implementa Pessoa) para a função
}
```
* `joao := Cliente{...}`: Criamos uma variável `joao` do tipo `Cliente` e a inicializamos.
* `joao.Cidade = "Brasília"`: Definimos a cidade de João. Lembre-se que `Cidade` é um campo da struct `Endereco` que está incorporada em `Cliente`.
* `Desativacao(joao)`: Em vez de chamar `joao.Desativar()` diretamente, chamamos a função `Desativacao` e passamos `joao` como argumento. Como `Cliente` tem o método `Desativar()`, ele automaticamente satisfaz a interface `Pessoa`, e, portanto, pode ser passado para `Desativacao`.

Este código ilustra como você pode organizar seus dados com structs, definir comportamentos com interfaces e criar métodos para operar em seus tipos, tudo isso de forma concisa e eficiente em Go.

---