package main

import (
	"fmt"
	"github.com/go-algorithms/Search/minmax"
)

func main() {
	numbers := []int{11, 59, 3, 2, 53, 17, 31, 7, 19, 67, 47, 13, 37, 61, 29, 43, 5, 41, 23}

	result := minmax.Minimum(numbers)

	fmt.Printf("Minimum number in %v is (%v)\n", numbers, result)

}