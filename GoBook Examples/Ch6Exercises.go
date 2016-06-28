package main

import "fmt"

func main()  {

  x := [4]int{1, 2, 3, 4}
  y := make([]int, 3)
  z := make([]int, 3, 9)
  a := [6]string{"a", "b", "c", "d", "e", "f"}
  b := []int{
    48,96,86,68,
    57,82,63,70,
    37,34,83,27,
    19,97, 9,17,
    }

  fmt.Println(x)
  fmt.Println(x[3])

  fmt.Println(y)
  fmt.Println(y[2])

  fmt.Println(z)
  fmt.Println(len(z))

  fmt.Println(a[2:5])

  var buff int = b[0]
  for _, value := range b {
    if value < buff {
      buff = value
    }
  }
  fmt.Println(buff)

}
