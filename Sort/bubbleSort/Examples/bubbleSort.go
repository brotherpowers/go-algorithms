package main

import (
	"fmt"

	"github.com/brotherpowers/go-algorithms/Sort/bubbleSort"
)

func main() {

	numbers := []int{11, 59, 3, 2, 53, 17, 31, 7, 19, 67, 47, 13, 37, 61, 29, 43, 5, 41, 23}

	fmt.Println("Array before sorting", numbers)

	// Performing bubbleSort
	fmt.Println("Array after sorting", bubbleSort.Sort(numbers))
}
