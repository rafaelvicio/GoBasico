package main

import (
  "fmt"
)

func main(){
  var numero1 int
  var numero2 int

  fmt.Print("Digite o primeiro número: ")
  fmt.Scanln(&numero1)

  fmt.Print("Digite o segundo número: ")
  fmt.Scanln(&numero2)

  fmt.Printf("%d + %d = %d \n", numero1, numero2, numero1+numero2)
  fmt.Printf("%d - %d = %d \n", numero1, numero2, numero1-numero2)
  fmt.Printf("%d x %d = %d \n", numero1, numero2, numero1*numero2)
  fmt.Printf("%d / %d = %d \n", numero1, numero2, numero1/numero2)
  fmt.Printf("%d %% %d = %d \n", numero1, numero2, numero1%numero2)

  incremento := numero1
  incremento += numero2

  decremento := numero1
  decremento -= numero2

  fmt.Printf("O incremento de %d com %d é %d \n", numero1, numero2, incremento)
  fmt.Printf("O drecremento de %d com %d é %d \n", numero1, numero2, decremento)
}
