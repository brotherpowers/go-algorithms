# Bubble Sort

Bubble sort is a simple sorting algorithm. It works by repeatedly stepping through the list to be sorted, comparing each pair of adjacent items and swapping them if they are in the wrong order. The pass through the list is repeated until no swaps are needed, which indicates that the list is sorted. Because it only uses comparisons to operate on elements, it is a comparison sort.


## The Code:

Here is an implementation of counting sort

```go
func Sort(inputArray []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i <= len(inputArray)-1; i++ {
			if inputArray[i-1] > inputArray[i] {
				inputArray[i], inputArray[i-1] = inputArray[i-1], inputArray[i]
				swapped = true
			}
		}
	}

	return inputArray
}
```

## Performance

* Worst case performance: **O(n2)**
* Best case performance: **O(n)**
* Average case performance: **O(n2)**
* Worst case space complexity: **O(n) total, O(1) auxiliary**


## See also

[Bubble Sort on Wikipedia](https://en.wikipedia.org/wiki/Bubble_sort)

*Written by [brotherpowers](https://www.brotherpowers.com/)*
