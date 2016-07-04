package main

import "fmt"

func max(x ...int) (int) {

  maximum := x[0]

  for i, value := range x {
    if value > maximum {
      maximum = x[i]
    }
  }
  return maximum
}

func main() {

  xs := []int{6,3,43,75,21}
  fmt.Println(max(xs...))
}
