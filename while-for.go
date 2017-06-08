package main

import(
	"fmt"
)

func main() {
	var numero int
	fmt.Print("Digite um número")
	fmt.Scanf("%d", &numero)

	i := 0
	for i <= 10 {
		fmt.Printf("%d x %d = %d \n", numero, i, numero*1)
		i++
	}
}
