# Radix Sort

Radix sort is a sorting algorithm that takes as input an array of integers and uses a sorting subroutine( that is often another efficient sorting algorith) to sort the integers by their radix, or rather their digit.  [Counting Sort](../countingSort/), and Bucket Sort are often times used as the subroutine for Radix Sort. For this implementation, we will use [Counting Sort](../countingSort/).

> **Note:** This implementation does not handle negative numbers.

## Example

```
Input [170, 45, 75, 90, 802, 24, 2, 66]
Output [2, 24, 45, 66, 75, 90, 170, 802]
```

### Step 1:
The first step in this algorithm is to define the digit or rather the "base" or radix that we will use to sort.
For this example we will `radix = 10`, since the integers we are working with in the example are of base 10.

### Step 2:
The next step is to simply iterate n times (where n is the number of digits in the largest integer in the input array), and upon each iteration perform a sorting subroutine on the current digit in question.

## The Code:

Here is an implementation of radix sort

```go
// Here we define our radix to be 10
const radix = 10

func countSort(arr []int, exp int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}

	// Create an array to store the count of each element
	count := make([]int, radix)

	// Array to hold the sorted values
	output := make([]int, length)

	for i := 0; i < length; i++ {
		count[(arr[i]/exp)%radix] += 1
	}

	// Change count[i] so that count[i] now contains
	// actual position of this digit in output[]
	for i := 1; i < radix; i++ {
		count[i] += count[i-1]
	}

	// Place the element in the final array as per the number of elements before it
	for i := length - 1; i >= 0; i-- {
		output[count[(arr[i]/exp)%radix]-1] = arr[i]
		count[(arr[i]/exp)%radix]--
	}
	return output
}

func Sort(arr []int) {

	// Find the maximum number to know number of digits
	max := minmax.Maximum(arr) + 1
	length := len(arr)

	for exp := 1; max/exp > 0; exp *= radix {
		output := countSort(arr, exp)

		for i := 0; i < length; i++ {
			arr[i] = output[i]
		}
	}
}
```

Let's take a look at our example input array.

The largest integer in our array is 802, and it has three digits (ones, tens, hundreds).  So our algorithm will iterate three times whilst performing some sorting algorithm on the digits of each integer.

* Iteration 1:  170, 90, 802, 2, 24, 45, 75, 66
* Iteration 2:  802, 2, 24, 45, 66, 170, 75, 90
* Iteration 3:  2, 24, 45, 66, 75, 90, 170, 802

## Performance

Let there be d digits in input integers. Radix Sort takes **O(d*(n+b))** time where b is the base for representing numbers, for example, for decimal system, b is 10. What is the value of d? If k is the maximum possible value, then d would be **O(logb(k))**. 

So overall time complexity is **O((n+b) * logb(k))**. Which looks more than the time complexity of comparison based sorting algorithms for a large k. Let us first limit k. Let k <= nc where c is a constant. In that case, the complexity becomes **O(nLogb(n))**. But it still doesn’t beat comparison based sorting algorithms.

See also [Wikipedia](https://en.wikipedia.org/wiki/Radix_Sort).

*Written by [brotherpowers](https://www.brotherpowers.com/)*
