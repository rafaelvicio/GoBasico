package main

import(
	"fmt"
)

func main(){
	amigos := make([]string, 3)
	nome := ""
	i := 0

	for nome != "q" {
		fmt.Print("Digite o nome d eum amigo ou 'q' para parar:")
		fmt.Scanf("%s", &nome)
		if nome != "q" {
			if i < 3 {
				amigos[i] = nome
			} else {
				amigos = append(amigos, nome)
			}
			i++
		}
	}
	fmt.Println(amigos)
	fmt.Printf("Você tem %d amigos \n", len(amigos))
	for _, nomeAmigo := range amigos {
		fmt.Printf("- %s \n", nomeAmigo)
	}
}
