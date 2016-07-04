package main

import "fmt"

func OddGen() func() uint  {
  i := uint(1)

  return func () (ret uint)  {
    ret = i
    i += 2

    return
  }

}


func main() {
  nextOdd := OddGen()

  fmt.Println(nextOdd())
  fmt.Println(nextOdd())
  fmt.Println(nextOdd())

}
