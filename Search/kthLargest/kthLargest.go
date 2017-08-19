package kthLargest

import (
	"math/rand"
	"time"
)

func randomPivot(arr *[]int, min, max int) int {
	// choose a random pivot index
	pivotIndex := randomNumber(min, max)
	// swap the element at pivot <---> max position
	swap(arr, pivotIndex, max)
	return (*arr)[max]
}

func randomPartition(arr *[]int, min, max int) int {
	// fmt.Println("before random pivot")
	pivot := randomPivot(arr, min, max)
	// fmt.Println("after random pivot")

	i := min
	for j := min; j < max; j++ {
		if (*arr)[j] <= pivot {
			swap(arr, i, j)
			i = i + 1
		}
	}
	swap(arr, i, max)
	return i
}

func randomSelect(arr *[]int, min, max, k int) int {
	if min < max {
		p := randomPartition(arr, min, max)

		if k == p {
			return (*arr)[p]
		} else if k < p {
			return randomSelect(arr, min, p-1, k)
		} else {
			return randomSelect(arr, p+1, max, k)
		}

	} else {
		return (*arr)[min]
	}
}

func KthLargest(inputArray []int, k int) int {
	length := len(inputArray)

	if length < 0 {
		panic("Array must have at-least one element")
	}
	if k >= length {
		k = length - 1
	}

	k = length - k - 1

	return randomSelect(&inputArray, 0, length-1, k)
}

func KthSmallest(inputArray []int, k int) int {
	length := len(inputArray)

	if length < 0 {
		panic("Array must have at-least one element")
	}

	if k >= length {
		k = length - 1
	}
	return randomSelect(&inputArray, 0, length-1, k)
}

// Returns a random integer in the range min...max, inclusive
func randomNumber(min, max int) int {
	rand.Seed(int64(time.Now().UnixNano()))
	return rand.Intn(max-min) + min
}

func swap(arr *[]int, i, j int) {
	temp := (*arr)[i]
	(*arr)[i] = (*arr)[j]
	(*arr)[j] = temp
}
