package main

import (
	"fmt"
	"github.com/go-algorithms/Search/kthLargest"
	"sort"
)

func main() {
	arr := []int{11, 59, 3, 2, 53, 17, 31, 7, 19, 67, 47, 13, 37, 61, 29, 43, 5, 41, 23}

	sort.Ints(arr)

	n := 5

	largest := kthLargest.KthLargest(arr, n)
	fmt.Println(n, "Largest number in", arr, "is", largest)

	smallest := kthLargest.KthSmallest(arr, n)
	fmt.Println(n, "Smallest number in", arr, "is", smallest)
}
