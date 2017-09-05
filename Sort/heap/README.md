# Heap Sort

Sorts an array from low to high using a heap.

A [heap](../heap/) is a partially sorted binary tree that is stored inside an array. The heap sort algorithm takes advantage of the structure of the heap to perform a fast sort.

To sort from lowest to highest, heap sort first converts the unsorted array to a max-heap, so that the first element in the array is the largest.

## An example

Let's say the array to sort is:

	[ 5, 13, 2, 25, 7, 17, 20, 8, 4 ]

This is first turned into a max-heap that looks like this:

![The max-heap](Images/MaxHeap.png)

The heap's internal array is then:

	[ 25, 13, 20, 8, 7, 17, 2, 5, 4 ]

That's hardly what you'd call sorted! But now the sorting process starts: we swap the first element (index *0*) with the last one at index *n-1*, to get:

	[ 4, 13, 20, 8, 7, 17, 2, 5, 25 ]
	  *                          *

Now the new root node, `4`, will be smaller than its children, so we fix up the max-heap up to element *n-2* using the *shift down* or "heapify" procedure. After repairing the heap, the new root is now the second-largest item in the array:

	[20, 13, 17, 8, 7, 4, 2, 5 | 25]

Important: When we fix the heap, we ignore the last item at index *n-1*. That now contains the array's maximum value, so it is in its final sorted place already. The `|` bar indicates where the sorted portion of the array begins. We'll leave that part of the array alone from now on.

Again, we swap the first element with the last one (this time at index *n-2*):

	[5, 13, 17, 8, 7, 4, 2, 20 | 25]
	 *                      *

And fix up the heap to make it valid max-heap again:

	[17, 13, 5, 8, 7, 4, 2 | 20, 25]

As you can see, the largest items are making their way to the back. We repeat this process until we arrive at the root node and then the whole array is sorted.


> **Note:** This process is very similar to [selection sort](../selection/), which repeatedly looks for the minimum item in the remainder of the array. Extracting the minimum or maximum value is what heaps are good at.

##The Code

```go
func heapify(arr []int, heapSize, root int) {
	largest := root
	left := 2*root + 1
	right := 2*root + 2

	if left < heapSize && arr[left] > arr[largest] {
		largest = left
	}

	if right < heapSize && arr[right] > arr[largest] {
		largest = right
	}

	if largest != root {
		arr[root], arr[largest] = arr[largest], arr[root]
		heapify(arr, heapSize, largest)
	}
}
```

A step-by-step explanation of how the code works:

1. Initialize largest as root.
2. If left child is larger than root then `largest = left`
3. If right child is larger than largest so far then `largest = right`
4. If largest is not root. Swap element at position `root` with `largest` `arr[root], arr[largest] = arr[largest], arr[root]` and recursively heapify the affected sub-tree

```go
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
```

## Performance

Performance of heap sort is **O(n lg n)** in best, worst, and average case. Because we modify the array directly, heap sort can be performed in-place. But it is not a stable sort: the relative order of identical elements is not preserved.

## See also

[Heap sort on Wikipedia](https://en.wikipedia.org/wiki/Heapsort)

*Written by [brotherpowers](https://www.brotherpowers.com/)*