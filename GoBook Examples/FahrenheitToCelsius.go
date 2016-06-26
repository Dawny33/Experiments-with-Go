package main

import "fmt"

func main() {

  fmt.Println("Enter temperature in F: ")

  var F float64
  fmt.Scanf("%f", &F)

  var Celsius float64
  Celsius = (F - 32) * 5/9

  fmt.Println("Temp. in Celsius", Celsius)
}
