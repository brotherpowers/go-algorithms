package kthLargest

import (
	"math/rand"
	"time"
)

func randomPivot(arr *[]int, low, high int) int {
	// choose a random pivot index
	pivotIndex := randomNumber(low, high)
	// swap the element at pivot <---> high position
	swap(arr, pivotIndex, high)
	return (*arr)[high]
}

func randomPartition(arr *[]int, low, high int) int {
	// fmt.Println("before random pivot")
	pivot := randomPivot(arr, low, high)
	// fmt.Println("after random pivot")

	i := low
	for j := low; j < high; j++ {
		if (*arr)[j] <= pivot {
			swap(arr, i, j)
			i = i + 1
		}
	}
	swap(arr, i, high)
	return i
}

func randomSelect(arr *[]int, low, high, k int) int {
	if low < high {
		p := randomPartition(arr, low, high)

		if k == p {
			return (*arr)[p]
		} else if k < p {
			return randomSelect(arr, low, p-1, k)
		} else {
			return randomSelect(arr, p+1, high, k)
		}

	} else {
		return (*arr)[low]
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

// Returns a random integer in the range low...high, inclusive
func randomNumber(low, high int) int {
	rand.Seed(int64(time.Now().UnixNano()))
	return rand.Intn(high-low) + low
}

func swap(arr *[]int, i, j int) {
	(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
}

/*
func swap(arr *[]int, i, j int) {
	temp := (*arr)[i]
	(*arr)[i] = (*arr)[j]
	(*arr)[j] = temp
}*/
