package main

const a = "Hello, world"

var (
	b bool    = true
	c int     = 10
	d string  = "nome"
	e float64 = 1.2
)

func main() {
	a := "X" //var a string
	a = "Z"  //sem : quando for reatribuir, := só na primeira vez
	println(a)
}
