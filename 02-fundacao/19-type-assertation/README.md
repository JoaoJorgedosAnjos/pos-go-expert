# Análise de Código: Interface Vazia e Type Assertion em Go

1.  **Declaração da Interface Vazia:**
    Primeiro, declaramos uma variável `minhaVar` do tipo `interface{}`, conhecida como interface vazia. Ela pode armazenar um valor de qualquer tipo. Neste caso, atribuímos a ela uma `string`.

    ```go
    var minhaVar interface{} = "Wesley Williams"
    ```

2.  **Type Assertion Direta (bem-sucedida):**
    Para usar o valor da interface, afirmamos que seu tipo é `string`. Como a afirmação está correta, o valor é extraído com sucesso e impresso na tela. O programa continua a execução normalmente.

    ```go
    println(minhaVar.(string))
    ```

3.  **Type Assertion Segura (com verificação):**
    Esta é a forma idiomática e segura de fazer uma afirmação de tipo. Tentamos converter `minhaVar` para `int`. A operação falha, mas não quebra o programa. A variável booleana `ok` retorna `false`, indicando o erro, e a variável `res` recebe o valor zero do tipo `int` (que é `0`).

    ```go
    res, ok := minhaVar.(int)
    fmt.Printf("O valor de res é %v e o resultado de ok é %v\n", res, ok)
    ```

4.  **Type Assertion Direta (que causa pânico):**
    Aqui, tentamos forçar a conversão para `int` de forma direta e insegura. Como o tipo real da variável é `string`, essa operação é inválida e causa um **pânico (panic)** em tempo de execução, o que interrompe o programa abruptamente.

    ```go
    res2 := minhaVar.(int)
    // A linha abaixo nunca será executada por causa do pânico.
    fmt.Printf("O valor de res2 é %v\n", res2)
    ```