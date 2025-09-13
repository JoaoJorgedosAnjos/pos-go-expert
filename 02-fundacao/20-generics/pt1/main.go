package main

type MyNumber int

// Utilizando Constraint
type Number interface {
	~int | float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

//Com generci T = tipo
/*
func Soma[T int | float64](m map[string]T) T{
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}
//Sem generic um tipo por função

func Soma(m map[string]int) int{
	var soma int
	for _, v := range m {
		soma += v
	}
	return soma
}
*/

func main() {
	m := map[string]int{"Wesley": 1000, "João": 2000, "Maria": 3000}
	m2 := map[string]float64{"Wesley": 1000.00, "João": 2000.00, "Maria": 3000.00}
	m3 := map[string]MyNumber{"Wesley": 1000, "João": 2000, "Maria": 3000}
	print(Soma(m))
	print(Soma(m2))
	print(Soma(m3))
}
