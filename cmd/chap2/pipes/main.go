package main

import (
	"fmt"
	"golang/chap2/pipes"
	"os"
)

func main() {
	fmt.Printf("string: number_of_occurrences\n\n")
	for key, value := range pipes.WordCount(os.Stdin) {
		fmt.Printf("%s: %d\n", key, value)
	}
}
