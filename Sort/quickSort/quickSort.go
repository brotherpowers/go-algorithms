package quickSort

import (
	"math/rand"
	"time"
)

func Swap(arr *[]int, i, j int) {
	(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
}

// Lomuto's partitioning scheme
func partitionLumoto(arr []int, low, high int) int {
	length := len(arr)

	if length == 0 {
		panic("Array size can not be 0")
	}

	if length == 1 {
		return -1
	}

	i := low
	pivot := arr[high]

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			Swap(&arr, i, j)
			i++
		}
	}

	Swap(&arr, i, high)
	return i
}

// Sort using Lomuto's partitioning
func SortLumoto(arr []int, low, high int) {
	if low < high {
		p := partitionLumoto(arr, low, high)
		SortLumoto(arr, low, p-1)
		SortLumoto(arr, p+1, high)
	}
}

// Hoare's partitioning scheme
func partitionHoare(arr []int, low, high int) int {
	length := len(arr)

	if length == 0 {
		panic("Array size is 0")
	}

	pivot := arr[low]
	i := low - 1
	j := high + 1

	for {
		for {
			j--

			if arr[j] <= pivot {
				break
			}
		}

		for {
			i++
			if arr[i] >= pivot {
				break
			}
		}

		if i < j {
			Swap(&arr, i, j)
		} else {
			return j
		}
	}
}

// Sort using Hoare's partitioning scheme
func SortHoare(arr []int, low, high int) {
	if low < high {
		p := partitionHoare(arr, low, high)
		SortHoare(arr, low, p)
		SortHoare(arr, p+1, high)
	}
}

// Quick Sort Random
func SortRandom(arr []int, low, high int) {
	if low < high {
		pivotIndex := random(low, high)
		arr[pivotIndex], arr[high] = arr[high], arr[pivotIndex]

		p := partitionLumoto(arr, low, high)
		SortRandom(arr, low, p-1)
		SortRandom(arr, p+1, high)
	}
}

// Returns a random int within range (low, high)
func random(low, high int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(high-low) + low
}

func partitionDutchFlag(arr []int, low, high, pivotIndex int) (int, int) {
	length := len(arr)

	if length == 0 {
		panic("Array size can not be zero")
	}

	pivot := arr[pivotIndex]
	smaller := low
	equal := low
	larger := high

	for equal <= larger {
		if arr[equal] < pivot {
			Swap(&arr, smaller, equal)
			smaller++
			equal++
		} else if arr[equal] == pivot {
			equal++
		} else {
			Swap(&arr, equal, larger)
			larger--
		}
	}

	return smaller, larger
}

func SortDutchFlag(arr []int, low, high int) {
	if low < high {
		pivotIndex := random(low, high)
		p, q := partitionDutchFlag(arr, low, high, pivotIndex)
		SortDutchFlag(arr, low, p-1)
		SortDutchFlag(arr, q+1, high)
	}
}
