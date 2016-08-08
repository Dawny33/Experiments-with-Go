package main

// This can be checked by running the program as follows:
// go run CommandLine_Args.go some args

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	// We start the loop from 1, as Args[0] gives this program's name
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
