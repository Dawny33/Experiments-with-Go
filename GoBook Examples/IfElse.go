package main

import "fmt"

func main()  {

  for i := 1; i <= 11; i++ {
    if i % 2 == 0 {
      fmt.Println(i, " is Even")
    } else {
      fmt.Println(i, " is Odd")
    }
  }
}
