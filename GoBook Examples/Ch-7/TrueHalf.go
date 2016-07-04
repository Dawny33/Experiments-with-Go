package main

import "fmt"

func truehalf(x int) (int, bool)  {

  if x%2 == 0 {
    return int(x/2), true
  } else {
    return int(x/2), false
  }
}



func main() {
  fmt.Println(truehalf(1))
  fmt.Println(truehalf(2))
}
