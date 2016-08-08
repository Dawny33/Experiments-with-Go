package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string

	// For loop with range always return index, element
	// So, if the index is not needed, one can use a blank identifier.
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
