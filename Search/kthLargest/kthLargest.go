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

func KthLargest(arr []int, k int) int {
	length := len(arr)

	if length < 0 {
		panic("Array must have at-least one element")
	}
	if k >= length {
		k = length - 1
	}

	k = length - k - 1

	return randomSelect(&arr, 0, length-1, k)
}

// Returns a random integer in the range low...high, inclusive
func randomNumber(low, high int) int {
	rand.Seed(int64(time.Now().UnixNano()))
	return rand.Intn(high-low) + low
}

func swap(arr *[]int, i, j int) {
	(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
}

// Find the Kth largest nnumber using Heap
func KthLargestUsingHeap(arr []int, k int) int {

	var heapify func(arr []int, heapSize, root int)
	heapify = func(arr []int, heapSize, root int) {
		largest := root
		left := 2*root + 1
		right := 2*root + 2

		// If left child is larger than root
		if left < heapSize && arr[left] > arr[largest] {
			largest = left
		}

		// If right child is larger than largest so far
		if right < heapSize && arr[right] > arr[largest] {
			largest = right
		}

		// If largest is not root
		if largest != root {
			arr[root], arr[largest] = arr[largest], arr[root]
			heapify(arr, heapSize, largest) // Recursively heapify the affected sub-tree
		}
	}

	length := len(arr)
	if length == 0 {
		panic("Array size can not be zero")
	}
	if length == 1 {
		return arr[0]
	}

	// Build heap (rearrange array)
	for i := length/2 - 1; i >= 0; i-- {
		heapify(arr, length, i)
	}
	for i := length - 1; i >= (length - 1 - k); i-- {
		arr[0], arr[i] = arr[i], arr[0] // Move current root to end
		heapify(arr, i, 0)              // call max heapify on the reduced heap
	}
	return arr[length-1-k]
}
