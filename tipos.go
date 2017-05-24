package main

import (
  "fmt"
)

func main() {
  fmt.Println("INTEIROS SEM SINAL")
  var u1 byte = 255
  var u2 uint16 = 33
  var u3 uint32 = 44
  var u4 uint64 = 55
  fmt.Println(u1)
  fmt.Println(u2)
  fmt.Println(u3)
  fmt.Println(u4)

  fmt.Println("INTEIROS COM SINAL")
  var i1 int8 = 127
  var i2 int16 = 1000
  var i3 int32 = 1001
  var i4 int64 = 1002
  fmt.Println(i1)
  fmt.Println(i2)
  fmt.Println(i3)
  fmt.Println(i4)

  var t1 uint = 1010
  var t2 int = 2000
  fmt.Println(t1)
  fmt.Println(t2)

  fmt.Println("PONTOS FLUTUANTES")
  var f1 float32 = 1
  var f2 float64 = 2
  var f3 complex64 = 3
  var f1 complex128 = 4
  fmt.Println(f1)
  fmt.Println(f2)
  fmt.Println(f3)
  fmt.Println(f4)

  fmt.Println("STRINGS")
  var s1 string = "Go"
  var s2 string = `Go`
  fmt.Println(s1)
  fmt.Println(s2)

  fmt.Println("MULTIPLAS DECLARAÇÔES")
  var nome, sobrenome string = "Go", "Lang"
  fmt.Println(nome + " " + sobrenome)

}
