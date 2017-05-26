package main

import (
  "fmt"
)

func main(){
  const constante string = "Go"
  fmt.Println(constante)

  constante = "Go Lang" // Aqui vai da erro

  const (
    primeiroNome = "Rafael"
    segundoNome = "Augusto"
  )

  fmt.Println(primeiroNome + " " + segundoNome)
}
