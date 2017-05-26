package main

import(
  "fmt"
  "strconv"
)

func main(){

  var texto string
  fmt.Print("Digite um número: ")
  fmt.Scanf("%s", &texto)

  // var numero int
  // numero, _ = strconv.Atoi(texto)

  numero, _ := strconv.ParseInt(texto, 10, 32)
  var conv = float64(numero)
  fmt.Println(numero)
  fmt.Println(conv)
}
