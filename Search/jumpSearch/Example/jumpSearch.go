package main

import (
	"fmt"

	"github.com/brotherpowers/Search/jumpSearch"
)

func main() {
	input := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	searchFor := 17

	fmt.Printf("Searching for %v in %v\n", searchFor, input)

	position := jumpSearch.Search(input, searchFor)

	fmt.Printf("Element found at position: %v\n", position)
}
