package main

import (
	"fmt"
	"github.com/go-algorithms/Search/nthLargest"
)

func main() {
	arr := []int{11, 59, 3, 2, 53, 17, 31, 7, 19, 67, 47, 13, 37, 61, 29, 43, 5, 41, 23}

	n := 5

	largest := nthLargest.NthLargest(arr, n)
	fmt.Println(n, "Largest number in", arr, "is", largest)

	smallest := nthLargest.NthSmallest(arr, n)
	fmt.Println(n, "Smallest number in", arr, "is", smallest)
}
