# Minimum | Maximum

Goal: Find the minimum|maximum object in an unsorted array.

## Maximum or minimum

We have an array of generic objects and we iterate over all the objects keeping track of the minimum/maximum element so far.

### An example

Let's say the we want to find the maximum value in the unsorted list `[ 8, 3, 9, 4, 6 ]`.

Pick the first number, `8`, and store it as the maximum element so far. 

Pick the next number from the list, `3`, and compare it to the current maximum. `3` is less than `8` so the maximum `8` does not change.

Pick the next number from the list, `9`, and compare it to the current maximum. `9` is greater than `8` so we store `9` as the maximum.

Repeat this process until the all elements in the list have been processed.

### The code

```go
func Minimum(inputArray []int) int {
	if len(inputArray) == 0 {
		panic("Array must have at-least one element")
	}

	min := inputArray[0]

	for _, value := range inputArray {
		if value < min {
			min = value
		}
	}
	
	return min
}


func Maximum(inputArray []int) int {
	if len(inputArray) == 0 {
		panic("Array must have at-least one element")
	}

	max := inputArray[0]

	for _, value := range inputArray {
		if value > max {
			max = value
		}
	}
	
	return max
}
```

```go
numbers := []int{11, 59, 3, 2, 53, 17, 31, 7, 19, 67, 47, 13, 37, 61, 29, 43, 5, 41, 23}
minmax.Minimum(array)   // This will return 2
minmax.Maximum(array)   // This will return 67
```

By picking elements in pairs and comparing their maximum and minimum with the running minimum and maximum we reduce the number of comparisons to 3 for every 2 elements.

## Performance

These algorithms run at **O(n)**. Each object in the array is compared with the running minimum/maximum so the time it takes is proportional to the array length.

*Written by [brotherpowers](https://www.brotherpowers.com/)*