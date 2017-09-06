# Counting Sort

Counting sort is an algorithm for sorting a collection of objects according to keys that are small integers. It operates by counting the number of objects that have each distinct key values, and using arithmetic on those counts to determine the positions of each key value in the output sequence.

## Example

To understand the algorithm let's walk through a small example.

Consider the array: `[ 10, 9, 8, 7, 1, 2, 7, 3 ]`

### Step 1:

The first step is to count the total number of occurrences for each item in the array. The output for the first step would be a new array that looks as follows:

```
Index 0 1 2 3 4 5 6 7 8 9 10
Count 0 1 1 1 0 0 0 2 1 1 1
```

Here is the code to accomplish this:

```go
	max := minmax.Maximum(arr)
	countArray := make([]int, max+1)
	for _, v := range arr {
		countArray[v] += 1
	}
```

> **Note:** `max` is the maximum number in the array is found by using [minmax](../../Search/minmax/).

### Step 2:

In this step the algorithm tries to determine the number of elements that are placed before each element. Since, you already know the total occurrences for each element you can use this information to your advantage. The way it works is to sum up the previous counts and store them at each index.

The count array would be as follows:

```
Index 0 1 2 3 4 5 6 7 8 9 10
Count 0 1 2 3 3 3 3 5 6 7 8
```

The code for step 2 is:

```go
	for i := 1; i < len(countArray); i++ {
		countArray[i] += countArray[i-1]
	}
```

### Step 3:

This is the last step in the algorithm. Each element in the original array is placed at the position defined by the output of step 2. For example, the number 10 would be placed at an index of 7 in the output array. Also, as you place the elements you need to reduce the count by 1 as those many elements are reduced from the array.

The final output would be:

```
Index  0 1 2 3 4 5 6 7
Output 1 2 3 7 7 8 9 10
```

Here is the code for this final step:

```go
	sortedArray := make([]int, length);
	for _, v := range arr {
		countArray[v] -= 1
		sortedArray[countArray[v]] = v
	}
```

## The Code:

Here is an implementation of counting sort

```go
	func Sort(arr []int) []int {
		length := len(arr)
		if length <= 1 {
			return arr
		}

		max := minmax.Maximum(arr)
		// Create an array to store the count of each element
		countArray := make([]int, max+1)
		for _, v := range arr {
			countArray[v] += 1
		}

		// Set each value to be the sum of the previous two values
		for i := 1; i < len(countArray); i++ {
			countArray[i] += countArray[i-1]
		}

		// Array to hold the sorted values
		sortedArray := make([]int, length)

		// Place the element in the final array as per the number of elements before it
		for _, v := range arr {
			countArray[v] -= 1
			sortedArray[countArray[v]] = v
		}
		return sortedArray
	}
```

> **Note:** This implementation doesn't handle `negative numbers`.

## Performance

The algorithm uses simple loops to sort a collection. Hence, the time to run the entire algorithm is **O(n+k)** where **O(n)** represents the loops that are required to initialize the output arrays and **O(k)** is the loop required to create the count array.

The algorithm uses arrays of length **n + 1** and **n**, so the total space required is **O(2n)**. Hence for collections where the keys are scattered in a dense area along the number line it can be space efficient.

## See also

[Quick sort on Wikipedia](https://en.wikipedia.org/wiki/Counting_sort)

*Written by [brotherpowers](https://www.brotherpowers.com/)*
