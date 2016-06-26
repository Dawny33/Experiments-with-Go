package main

import "fmt"

func main() {

  fmt.Print("Enter the height in feet: ")

  var feet float64
  fmt.Scanf("%f", &feet)

  var metres float64
  metres = feet * 0.3048

  fmt.Println("The height in metres is: ", metres)
  
}
