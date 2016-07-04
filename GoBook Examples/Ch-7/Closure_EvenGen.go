package main

import "fmt"

func EvenGen() func() uint {
  i := uint(0)

  return func () (ret uint)  {
    ret = i
    i += 2
    return

  }
}

func main() {
  nextEven := EvenGen()

  fmt.Println(nextEven())
  fmt.Println(nextEven())
  fmt.Println(nextEven())
}
