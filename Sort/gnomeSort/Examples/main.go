package main

import (
	"fmt"

	"github.com/brotherpowers/go-algorithms/Sort/gnomeSort"
)

func main() {

	// Sort using Gnome Sort
	arr1 := []int{10, 0, 3, 9, 2, 14, 26, 27, 1, 5, 8, -1, 8}
	fmt.Println("Original array for Lumoto quickSort.Sort", arr1)
	gnomeSort.Sort(arr1)
	fmt.Println("Array after sorting", arr1)

}
