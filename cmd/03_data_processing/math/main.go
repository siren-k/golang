package main

import (
	"fmt"
	math2 "golang/internal/03_data_processing/math"
)

func main() {
	math2.Examples()

	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", math2.Fib(i))
	}
	fmt.Println()
}
