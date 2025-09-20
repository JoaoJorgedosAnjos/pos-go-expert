package main

import "fmt"

func main() {

    // --- 1. O 'if' Simples ---
    /*
        O 'if' simples executa um bloco de código apenas se a condição for verdadeira.
    */
    idade := 20
    if idade >= 18 {
        fmt.Println("Você é maior de idade e pode dirigir.")
    }

    // --- 2. O 'if' e 'else' ---
    /*
        O 'if-else' permite executar um bloco de código
        se a condição for verdadeira e um bloco diferente
        se a condição for falsa.
    */
    nota := 7.5
    if nota >= 7.0 {
        fmt.Println("Parabéns, você foi aprovado!")
    } else {
        fmt.Println("Que pena, você não foi aprovado.")
    }

    // --- 3. O 'if-else if-else' ---
    /*
        Usamos 'else if' para encadear múltiplas condições.
        A primeira condição que for verdadeira é a que será executada,
        e as outras serão ignoradas.
    */
    pontuacao := 85

    if pontuacao >= 90 {
        fmt.Println("Sua nota é A.")
    } else if pontuacao >= 80 {
        fmt.Println("Sua nota é B.")
    } else if pontuacao >= 70 {
        fmt.Println("Sua nota é C.")
    } else {
        fmt.Println("Sua nota é D.")
    }

    // --- 4. O 'if' com Declaração Curta (Short Statement) ---
    /*
        Este é um padrão muito comum em Go. Você pode declarar uma
        variável e usá-la na condição, tudo na mesma linha. A variável
        'num' só existe dentro deste bloco 'if'.
    */
    if num := 10; num > 5 {
        fmt.Println("O número", num, "é maior que 5.")
    } else {
        fmt.Println("O número", num, "não é maior que 5.")
    }

    // A linha abaixo causaria um erro de compilação porque 'num'
    // não existe fora do bloco 'if' acima.
    // fmt.Println(num)
}