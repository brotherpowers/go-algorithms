package main

import (
	"fmt"

	"github.com/brotherpowers/go-algorithms/Sort/radixSort"
)

func main() {

	numbers := []int{11, 59, 3, 2, 53, 17, 31, 7, 19, 67, 47, 13, 37, 61, 29, 43, 5, 41, 23}

	fmt.Println("Array before sorting", numbers)

	// Performing Radix sort
	radixSort.Sort(numbers)
	fmt.Println("Array after sorting", numbers)
}
