package countingSort

import (
	"github.com/brotherpowers/go-algorithms/Search/minmax"
)

func Sort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}

	max := minmax.Maximum(arr)
	// Create an array to store the count of each element
	countArray := make([]int, max+1)
	for _, v := range arr {
		countArray[v] += 1
	}

	// Set each value to be the sum of the previous two values
	for i := 1; i < len(countArray); i++ {
		countArray[i] += countArray[i-1]
	}

	// Array to hold the sorted values
	sortedArray := make([]int, length)

	// Place the element in the final array as per the number of elements before it
	for _, v := range arr {
		countArray[v] -= 1
		sortedArray[countArray[v]] = v
	}
	return sortedArray
}
