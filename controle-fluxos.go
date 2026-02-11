package main

import "fmt"

func main() {
	idade := 70
	if idade >= 65 {
		fmt.Println("Você é idoso.")	
	}else if idade >= 18 {
		fmt.Println("Você é maior de idade.")
	} else {
		fmt.Println("Você é menor de idade.")
	}

	diaDaSemana := "sábado"
	switch diaDaSemana {
	case "segunda":
		fmt.Println("Segunda-feira")
	case "terça":
		fmt.Println("Terça-feira")
	case "quarta":
		fmt.Println("Quarta-feira")
	case "quinta":
		fmt.Println("Quinta-feira")
	case "sexta":
		fmt.Println("Sexta-feira")
	case "sábado", "domingo":
		fmt.Println("Fim de semana")	
	default:
		fmt.Println("Dia inválido")		
	}
}