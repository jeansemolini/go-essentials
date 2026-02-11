package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	for {
		fmt.Println("Loop infinito")
		break
	}

	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
}