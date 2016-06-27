package main

import "fmt"

func main() {

    slice1 := []int{1,2,3,4}
    slice2 := make([]int, 2)

    copy(slice1, slice2)
    
    fmt.Println(slice1, slice2)
}
