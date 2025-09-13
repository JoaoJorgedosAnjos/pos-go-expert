package main

import "fmt"

func main()  {
	var x interface{} = 10
	var y interface{} = "Hello, world"

	showType(x)
	showType(y)
}

func showType(t interface{}){
	fmt.Printf("O tipo da variável é %T e o valor é %v\n",t,t)
}

/*go run main.go 
O tipo da variável é int e o valor é 10
O tipo da variável é string e o valor é Hello, world*/