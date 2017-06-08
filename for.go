package main

import(
	"fmt"
)

func main() {
	var numero int
	fmt.Print("Digite um número")
	fmt.Scanf("%d", &numero)

	for i := 0; i <= 10; i++ {
		fmt.Printf("%d x %d = %d \n", numero, i, numero*i)
	}
}