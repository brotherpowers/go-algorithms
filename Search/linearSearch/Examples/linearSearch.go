package main

import (
	"fmt"

	"github.com/go-algorithms/Search/linearSearch"
)

func main() {
	arr := []int{1, 2, 7, 8, 3, 4, 6, 9, 5}
	searchFor := 4

	fmt.Printf("Linear search: %v in %v\n", searchFor, arr)

	// Running the algo
	position := linearSearch.Search(arr, searchFor)

	fmt.Printf("Found At Position: %d", position)
}
