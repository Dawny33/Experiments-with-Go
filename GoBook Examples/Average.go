package main

import "fmt"

func main() {
  var x [5]float64
  x[0] = 34
  x[1] = 42
  x[2] = 45
  x[3] = 53
  x[4] = 97

  var total float64 = 0

  for _, value := range x {
    total += value
    }

  fmt.Println(total/ float64(len(x)) )
}
