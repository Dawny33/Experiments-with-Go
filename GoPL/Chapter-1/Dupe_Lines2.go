package main

import (
	"bufio"
	"fmt"
	"os"
)

// Takes an files and returns the count of each line
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()]++
	}
}

func main() {
	counts := make(map[string]int) //A map for lines and their counts
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg) //File opening event
			if err != nil {
				fmt.Print("Error:", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("Line: %s \t Count: %d \n", line, n)
		}
	}
}
