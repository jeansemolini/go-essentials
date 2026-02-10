package main

import "fmt"

func main() {
	a := 10
	b := 5
	soma := a + b
	subtracao := a - b
	multiplicacao := a * b
	divisao := a / b
	resto := a % b

	fmt.Println("Soma:", soma)
	fmt.Println("Subtração:", subtracao)
	fmt.Println("Multiplicação:", multiplicacao)
	fmt.Println("Divisão:", divisao)
	fmt.Println("Resto:", resto)
}