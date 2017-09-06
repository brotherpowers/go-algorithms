package radixSort

import (
	"github.com/brotherpowers/go-algorithms/Search/minmax"
)

func countSort(arr []int, exp int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}

	// Create an array to store the count of each element
	count := make([]int, length)

	// Array to hold the sorted values
	output := make([]int, length)

	for _, v := range arr {
		count[(v/exp)%10] += 1
	}

	// Change count[i] so that count[i] now contains
	// actual position of this digit in output[]
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	// Place the element in the final array as per the number of elements before it
	for i := length - 1; i >= 0; i-- {
		output[count[(arr[i]/exp)%10]-1] = arr[i]
		count[(arr[i]/exp)%10]--
	}
	return output
}

func Sort(arr []int) {

	// Find the maximum number to know number of digits
	max := minmax.Maximum(arr)
	length := len(arr)

	for exp := 1; max/exp > 0; exp *= 10 {
		output := countSort(arr, exp)

		for i := 0; i < length; i++ {
			arr[i] = output[i]
		}
	}
}
