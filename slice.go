package main

import "fmt"

func main() {
	numeros := []int{10, 20, 30}
	numeros = append(numeros, 40, 50)
	fmt.Println(numeros)
}