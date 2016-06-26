package main

import "fmt"

func main()  {
  fmt.Print("Enter a number: ")

  var input float64
  fmt.Scanf("%f", &input)

  output := input * 6
  fmt.Println(output)
}
