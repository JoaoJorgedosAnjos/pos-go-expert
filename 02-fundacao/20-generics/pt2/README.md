## A Interface `comparable` em Go

A interface `comparable` é uma **constraint pré-definida** do Go que permite apenas tipos que podem ser comparados usando os operadores `==` e `!=`.

### O que é `comparable`

```go
// Esta é a definição interna do Go (simplificada)
type comparable interface {
    // Tipos que suportam == e !=
}
```

É uma interface especial que o compilador entende automaticamente. Você **não precisa defini-la**, ela já existe no Go.

### Tipos que são `comparable`

✅ **Permitidos:**
- Tipos básicos: `int`, `float64`, `string`, `bool`
- Ponteiros: `*int`, `*string`
- Arrays de tipos comparáveis: `[3]int`
- Structs com campos comparáveis
- Interfaces
- Canais

❌ **NÃO permitidos:**
- Slices: `[]int`
- Maps: `map[string]int`
- Funções: `func()`

### Exemplo Prático

```go
package main

import "fmt"

// Função genérica que aceita qualquer tipo comparável
func Compara[T comparable](a, b T) bool {
    return a == b
}

// Função para encontrar item em slice
func Contem[T comparable](slice []T, item T) bool {
    for _, v := range slice {
        if v == item {
            return true
        }
    }
    return false
}

func main() {
    // Com inteiros
    fmt.Println(Compara(10, 10))     // true
    fmt.Println(Compara(10, 20))     // false
    
    // Com strings
    fmt.Println(Compara("Go", "Go")) // true
    fmt.Println(Compara("Go", "Rust")) // false
    
    // Com floats
    fmt.Println(Compara(3.14, 3.14)) // true
    
    // Usando em slice
    nums := []int{1, 2, 3, 4, 5}
    fmt.Println(Contem(nums, 3))     // true
    fmt.Println(Contem(nums, 9))     // false
    
    // Com strings
    palavras := []string{"Go", "Python", "Java"}
    fmt.Println(Contem(palavras, "Go"))     // true
    fmt.Println(Contem(palavras, "C++"))    // false
}
```

### Por que `Compara(10, 10.00)` funciona?

```go
func Compara[T comparable](a T, b T) bool {
    if a == b {
        return true
    }
    return false
}

// Esta chamada funciona:
println(Compara(10, 10.00)) // true
```

**Motivo:** Inferência de tipos com literais não tipados:

1. `10` é um literal inteiro **não tipado**
2. `10.00` é um literal float **não tipado**
3. O compilador precisa escolher um tipo `T` comum
4. Ambos podem ser representados como `float64`
5. Go escolhe `float64` como tipo comum
6. `10` vira `10.0`, `10.00` vira `10.0`
7. `10.0 == 10.0` → `true`

### Exemplo que DARIA erro:

```go
var a int = 10
var b float64 = 10.0
println(Compara(a, b)) // ERRO: tipos diferentes já definidos
```

### Casos de Uso Comuns

```go
// Busca genérica
func Encontrar[T comparable](slice []T, target T) int {
    for i, v := range slice {
        if v == target {
            return i
        }
    }
    return -1
}

// Remoção de duplicatas
func RemoverDuplicatas[T comparable](slice []T) []T {
    seen := make(map[T]bool)
    result := []T{}
    
    for _, v := range slice {
        if !seen[v] {
            seen[v] = true
            result = append(result, v)
        }
    }
    return result
}

// Uso
nums := []int{1, 2, 2, 3, 3, 4}
fmt.Println(RemoverDuplicatas(nums)) // [1 2 3 4]
```

### Resumo

- **`comparable`** = tipos que suportam `==` e `!=`
- **Uso principal:** funções que precisam comparar valores
- **Vantagem:** uma função funciona com qualquer tipo comparável
- **Cuidado:** literais não tipados podem ter inferência inesperada