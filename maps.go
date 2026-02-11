package main

import "fmt"

func main() {
	// idades := make(map[string]int)

	idades := map[string]int{
		"João": 30,
		"Maria": 25,
		"Pedro": 35,
	}

	idadeJoao := idades["João"]
	fmt.Println("Idade do João:", idadeJoao)

	idades["Ana"] = 28

	fmt.Println(idades)

	idade, ok := idades["Carlos"]
	if ok {
		fmt.Println("Idade do Carlos:", idade)
	} else {
		fmt.Println("Carlos não encontrado no mapa.")
	}

	delete(idades, "Pedro")
	fmt.Println(idades)

	for nome, idade := range idades {
		fmt.Println("Nome:", nome, "Idade:", idade)
	}
}