package main

import "fmt"

func swap(x, y *int) {
	
	tmpx := *x
	tmpy := *y

	*x = tmpy
	*y = tmpx

}

func main() {
	
	x := 3
	y := 4

	fmt.Println(x, y)
	
	swap(&x, &y)

	fmt.Println(x, y)
}