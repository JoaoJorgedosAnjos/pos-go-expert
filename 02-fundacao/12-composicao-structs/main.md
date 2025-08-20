### Entendendo a Composição de `structs` em Go

Este código demonstra um conceito fundamental em Go: a **composição de `structs`**. Em vez de usar herança (que Go não tem), você pode incorporar uma `struct` dentro de outra para reutilizar campos e comportamentos.

#### 1. Definição da `struct Endereco`

Começamos definindo uma `struct` simples para representar um endereço:

```go
type Endereco struct {
    Logradouro string
    Numero     int
    Cidade     string
    Estado     string
}
```

* `type Endereco struct { ... }`: Declara um novo tipo chamado `Endereco`.
* Ela contém quatro campos, todos com tipos básicos (`string` e `int`), para armazenar informações de endereço.

---

#### 2. Definição da `struct Cliente`: As Duas Formas de Composição

Aqui é onde a mágica (e a sua dúvida) acontece. A `struct Cliente` pode incorporar `Endereco` de duas maneiras principais:

```go
type Cliente struct {
    Nome   string
    Idade  int
    Ativo  bool
    Endereco      // Forma 1: Embedding (Incorporação Anônima)
    //ou
    //Adress Endereco // Forma 2: Composição Explícita (Campo Nomeado)
}
```

Vamos entender a diferença entre `Endereco` e `Adress Endereco`:

##### Forma 1: Embedding (Incorporação Anônima) - `Endereco`

Quando você simplesmente lista o **nome do tipo** da `struct` incorporada (neste caso, `Endereco`) sem um nome de campo explícito, isso é chamado de **embedding** ou **incorporação anônima**.

* **Sintaxe:** `Endereco` (apenas o nome do tipo).
* **Acesso aos Campos:** Os campos da `struct` incorporada (`Logradouro`, `Numero`, `Cidade`, `Estado`) são **promovidos** diretamente para a `struct` externa (`Cliente`). Isso significa que você pode acessá-los diretamente através de uma instância de `Cliente`, como se fossem campos próprios de `Cliente`.
    ```go
    joao.Cidade = "Brasília" // Acessa Cidade diretamente via joao
    // É o mesmo que joao.Endereco.Cidade = "Brasília" (Go faz isso por debaixo dos panos)
    ```
* **Vantagens:**
    * **Conveniência:** Acesso mais curto e direto aos campos da `struct` incorporada.
    * **Polimorfismo (com interfaces):** Facilita a satisfação de interfaces, pois os métodos da `struct` incorporada também são promovidos.
* **Desvantagens:**
    * Pode haver ambiguidade se a `struct` externa e a `struct` incorporada tiverem campos com o mesmo nome (nesse caso, você precisaria usar o nome completo `joao.Endereco.CampoDuplicado`).
    * Pode não ser tão explícito sobre a relação de "ter um" se você não estiver familiarizado com o padrão de embedding.

##### Forma 2: Composição Explícita (Campo Nomeado) - `Adress Endereco`

Quando você dá um **nome explícito** ao campo que é do tipo da `struct` incorporada (neste caso, `Adress`), isso é uma **composição explícita** ou um campo nomeado.

* **Sintaxe:** `Adress Endereco` (nome do campo seguido pelo tipo).
* **Acesso aos Campos:** Para acessar os campos da `Endereco`, você deve fazer isso através do nome do campo que a contém (`Adress`).
    ```go
    joao.Adress.Cidade = "Brasília" // Acessa Cidade através do campo Adress
    ```
* **Vantagens:**
    * **Clareza:** É muito explícito que `Cliente` "tem um" `Adress` que é do tipo `Endereco`.
    * **Evita Ambiguidade:** Não há risco de conflito de nomes de campo, pois você sempre especifica o caminho completo.
* **Desvantagens:**
    * O acesso aos campos é um pouco mais longo.

---

#### 3. A Função `main`: Colocando em Prática

Agora, vamos ver como essas definições se manifestam no código `main`:

```go
func main() {
    joao := Cliente{
        Nome:  "João",
        Idade: 21,
        Ativo: true,
        // Ao criar a struct Cliente, se Endereco for incorporada anonimamente,
        // os campos de Endereco podem ser inicializados aqui diretamente,
        // ou você pode inicializar a struct Endereco separadamente.
        // Exemplo:
        // Endereco: Endereco{
        //    Logradouro: "Rua Exemplo",
        //    Numero: 123,
        //    Cidade: "São Paulo",
        //    Estado: "SP",
        // },
    }

    joao.Ativo = false
    // joao.Adress.Cidade = "Brasília" // Se você usasse 'Adress Endereco'
    joao.Cidade = "Brasília"       // Funciona com 'Endereco' (embedding)
    // ou
    // joao.Endereco.Cidade = "Brasília" // Também funciona com 'Endereco' (embedding)
    // Go "promove" os campos, mas você ainda pode referenciá-los pelo nome do tipo incorporado.

    fmt.Println(joao.Nome) // Imprime "João"
    fmt.Println(joao.Cidade) // Imprime "Brasília" (acessado via campo promovido)
    fmt.Println(joao.Endereco.Cidade) // Também imprime "Brasília"
}
```

* **Criação da Instância:** Uma variável `joao` do tipo `Cliente` é criada e inicializada. Note que, na inicialização, você não precisa se preocupar com os campos de `Endereco` se estiver usando o embedding, a menos que queira inicializá-los com valores específicos (como mostrado nos comentários do código).
* **Modificando Campos:**
    * `joao.Ativo = false`: Altera o campo `Ativo` diretamente na `struct Cliente`.
    * `joao.Cidade = "Brasília"`: Esta linha funciona porque `Cidade` é um campo **promovido** da `Endereco` incorporada. Go permite que você acesse `joao.Cidade` diretamente, mesmo que `Cidade` pertença à `Endereco` internamente.
    * `joao.Endereco.Cidade = "Brasília"`: Esta linha também funciona e é mais explícita. Mesmo com o embedding, você **ainda pode** acessar os campos através do nome do tipo incorporado (`Endereco`), caso queira ser mais claro ou evitar ambiguidades.

---

### Qual usar: `Endereco` (Embedding) ou `Adress Endereco` (Campo Nomeado)?

A escolha depende da **clareza** e da **semântica** que você quer transmitir:

* Use **`Endereco` (embedding)** quando a `struct` externa **é aprimorada por** ou **incorpora diretamente** as características da `struct` interna, e você quer acesso direto e conciso aos campos dela. É comum para comportamentos compartilhados ou características que são "parte integrante" da `struct` principal.
    * Exemplo: Um `Carro` tem um `Motor`. O `Motor` é parte fundamental do `Carro`. Acessar `carro.Potencia` é direto.

* Use **`Adress Endereco` (campo nomeado)** quando a `struct` externa **tem um** relacionamento explícito com a `struct` interna, e você quer que essa relação seja óbvia no código. É útil quando a `struct` interna representa uma propriedade distinta ou um papel claro.
    * Exemplo: Um `Cliente` tem um `Endereco`. É claro que `Cliente` e `Endereco` são entidades separadas, e o cliente simplesmente possui um endereço. Acessar `cliente.Endereco.Rua` é mais descritivo.

No seu exemplo com `Cliente` e `Endereco`, ambas as abordagens são válidas. A escolha entre `joao.Cidade` e `joao.Endereco.Cidade` é uma questão de preferência pessoal e do nível de clareza que você deseja no código. O Go "promove" o campo para você, tornando o acesso direto possível com o embedding.

A composição é uma ferramenta muito poderosa em Go para construir estruturas de dados complexas e reutilizáveis sem a complexidade da herança.