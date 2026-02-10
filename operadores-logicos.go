package main

import "fmt"

func main() {
	a := true
	b := false
	eLogico := a && b
	ouLogico := a || b
	naoLogico := !a

	fmt.Println("E lógico (AND):", eLogico)
	fmt.Println("Ou lógico (OR):", ouLogico)
	fmt.Println("Não lógico (NOT):", naoLogico)
}