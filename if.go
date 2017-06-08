package main

import (
	"fmt"
)

func main(){

	// if

	var idade int
	fmt.Print("Digite a sua idade: ")
	fmt.Scanf("%d", &idade)
	if idade >= 18 {
		fmt.Println("Você é maior de idade")
	} else {
		fmt.Println("Você é menor de idade")
	}

	// else if

	var numero int
	fmt.Print("Digite um número: ")
	fmt.Scanf("%d", &numero)
	if numero < 0 {
		fmt.Println("Número negativo")
	} else if numero >= 0 && numero <=9 {
		fmt.Println("Número positivo")
	} else {
		fmt.Println("Número positivo")
	}
}