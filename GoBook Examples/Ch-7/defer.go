package main

import "fmt"

func first()  {
  fmt.Println("first")
}

func second()  {
  fmt.Println("second")
}

func main() {
  defer second()
  first()
}
