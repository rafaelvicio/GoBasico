package main

import (
	"fmt"
)

func main(){
	var mes int 
	fmt.Print("Digite o número do mês: ")
	fmt.Scanf("%d", &mes)

	switch  mes {
	case 1:
		fmt.Println("Esse mês é Janeiro")
	case 2:
		fmt.Println("Esse mês e Fevereiro")
	case 3:
		fmt.Println("Esse mês e março")
	default:
		fmt.Println("Mês inválido")
	}

	// conjuntos

	switch  mes  {
	case 1, 2, 3:
		fmt.Println("Primeiro trimestre")
	case 4, 5, 6:
		fmt.Println("Segundo trimestre")
	}
}
