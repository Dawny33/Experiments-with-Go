package main

import "fmt"

func sum(args ...int) int  {
  sum := 0
  for _,value := range args {
    sum += value
  }
  return sum
}

func main() {
  xs := []int{1,2,3,4,5}
  fmt.Println(sum(xs...))
}
