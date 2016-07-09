package main

import "fmt"

func zero(xPtr *int)  {

  *xPtr = 0
}

func main() {
  x := new(int)

  zero(x)

  fmt.Println(*x)
}
