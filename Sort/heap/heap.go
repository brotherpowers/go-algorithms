package heap

// To heapify a subtree rooted with node root which is
// an index in arr[]. heapSize is size of heap
func heapify(arr []int, heapSize, root int) {
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

		// Recursively heapify the affected sub-tree
		heapify(arr, heapSize, largest)
	}
}

func Sort(arr []int) {
	length := len(arr)

	if length <= 1 {
		return
	}

	// Build heap (rearrange array)
	for i := length/2 - 1; i >= 0; i-- {
		heapify(arr, length, i)
	}

	// One by one extract an element from heap
	for i := length - 1; i >= 0; i-- {
		// Move current root to end
		arr[0], arr[i] = arr[i], arr[0]

		// call max heapify on the reduced heap
		heapify(arr, i, 0)
	}
}
