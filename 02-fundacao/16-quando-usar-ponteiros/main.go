package main

// apontamento do endereço do valor
func soma(a, b *int) int {
	*a = 50
	*b = 50
	return *a + *b
}

func main() {
	var1 := 10
	var2 := 20
	soma(&var1, &var2)// caso não invocar a func var1 e 2 não vão ter os valores alterados
	println(var1)
	println(var2)
}

/*
Copia do valor
func soma(a, b int) int {
	a = 50
	return  a + b
}

func main() {
	var1 := 10
	var2 := 20
	soma(var1,var2)
	println(var1) //aqui retorna 10, lá em cima retorna 50
}
*/
