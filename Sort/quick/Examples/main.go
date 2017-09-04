package main

import (
	"fmt"

	"github.com/brotherpowers/go-algorithms/Sort/quick"
)

func main() {
	arr1 := []int{10, 0, 3, 9, 2, 14, 26, 27, 1, 5, 8, -1, 8}
	fmt.Println("Original array for Lumoto Quick Sort", arr1)
	quick.SortLumoto(arr1, 0, len(arr1)-1)
	fmt.Println("Array after sorting", arr1)

	arr2 := []int{2, 56, 4, 77, 26, 98, 55, 3, 9, 2, 14, -1, 8}
	fmt.Println("Original array for Hoare Quick Sort", arr2)
	quick.SortHoare(arr2, 0, len(arr2)-1)
	fmt.Println("Array after sorting", arr2)

	arr3 := []int{2, 56, 4, 77, 26, 98, 55, 3, 9, 2, 14, -1, 8}
	fmt.Println("Original array for Random Quick Sort", arr3)
	quick.SortRandom(arr3, 0, len(arr3)-1)
	fmt.Println("Array after sorting", arr3)
}
