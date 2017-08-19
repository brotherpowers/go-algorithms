package main

import (
	"fmt"
	"github.com/go-algorithms/Search/countOccurrences"
)

func main() {
	numbers := []int{0, 1, 1, 3, 3, 3, 3, 6, 8, 10, 11, 11}
	searchTerm := 3

	result := countOccurrences.Count(numbers, searchTerm)

	if result == 0 {
		fmt.Printf("Number (%v) is not found in %v\n", searchTerm, numbers)
	} else {
		fmt.Printf("Total occurrences for number (%v) is %v\n in %v\n", searchTerm, result, numbers)
	}
}
