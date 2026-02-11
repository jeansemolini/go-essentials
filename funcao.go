package main

import "fmt"

func soma(a int, b int) int {
	return a + b
}

func dividir(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("divisão por zero não é permitida")
	}
	return a / b, nil
}

func executar(f func(int) int, x int) int {
	return f(x)
}

func main() {
	resultado := soma(10, 5)
	fmt.Println("Resultado da soma:", resultado)
	
	divisao, err := dividir(10, 0)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado da divisão:", divisao)
	}

	dobrar := func(x int) int {
		return x * 2
	}
	resultadoFuncao := executar(dobrar, 5)
	fmt.Println("Resultado da função de ordem superior:", resultadoFuncao)
}