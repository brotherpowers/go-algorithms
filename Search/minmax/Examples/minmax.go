package main

import (
	"fmt"

	"github.com/brotherpowers/go-algorithms/Search/minmax"
)

func main() {
	numbers := []int{11, 59, 3, 2, 53, 17, 31, 7, 19, 67, 47, 13, 37, 61, 29, 43, 5, 41, 23}

	max := minmax.Maximum(numbers)
	min := minmax.Minimum(numbers)

	fmt.Printf("Maximum number in %v is (%v)\n", numbers, max)
	fmt.Printf("Minimum number in %v is (%v)\n", numbers, min)

}
