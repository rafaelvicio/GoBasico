package main

import (
	"fmt"
	"bufio"
	"os"
)
func main(){
	mapaCursso := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	curso := ""
	for curso != "q" {
		var cargaHoraria int
		fmt.Print("Digite um curso ou digite q para sair: ")
		scanner.Scan()
		curso = scanner.Text()
		if curso != "q" {
			fmt.Print("Digite a carga horaria do curso: ")
			fmt.Scanf("%d", &cargaHoraria)
			mapaCursso[curso] = cargaHoraria
		}
	}
	fmt.Println("Cursos registrados: ")
	for nomeCurso, cargaHoraria := range mapaCursso {
		fmt.Print("- %s: %d \n ", nomeCurso, cargaHoraria)
	}
}
