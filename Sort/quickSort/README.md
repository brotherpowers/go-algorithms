# Quicksort

Goal: Sort an array from low to high (or high to low).

Quicksort (sometimes called *partition-exchange sort*) is an efficient sorting algorithm, serving as a systematic method for placing the elements of an array in order. Developed by Tony Hoare in 1959 and published in 1961, it is still a commonly used algorithm for sorting. When implemented well, it can be about *two or three times faster* than its main competitors, [merge sort](../merge/)) sort and [heap sort](../heap/).

##Algorithm
Quicksort is a divide and conquer algorithm. Quicksort first divides a large array into two smaller sub-arrays: the low elements and the high elements. Quicksort can then recursively sort the sub-arrays.

The steps are:

1. Pick an element, called a pivot, from the array.
2. Partitioning: reorder the array so that all elements with values less than the pivot come before the pivot, while all elements with values greater than the pivot come after it (equal values can go either way). After this partitioning, the pivot is in its final position. This is called the partition operation.
3. Recursively apply the above steps to the sub-array of elements with smaller values and separately to the sub-array of elements with greater values.
The base case of the recursion is arrays of size zero or one, which never need to be sorted.

The pivot selection and partitioning steps can be done in several different ways; the choice of specific implementation schemes greatly affects the algorithm's performance.

## Partitioning

Dividing the array around the pivot is called *partitioning* and there are a few different partitioning schemes.

If the array is,

	[ 10, 0, 3, 9, 2, 14, 8, 27, 1, 5, 8, -1, 26 ]

and we choose the middle element `8` as a pivot then after partitioning the array will look like this:

	[ 0, 3, 2, 1, 5, -1, 8, 8, 10, 9, 14, 27, 26 ]
	  -----------------        -----------------
	  all elements < 8         all elements > 8

The key thing to realize is that after partitioning the pivot element is in its final sorted place already. The rest of the numbers are not sorted yet, they are simply partitioned around the pivot value. Quicksort partitions the array many times over, until all the values are in their final places.

There is no guarantee that partitioning keeps the elements in the same relative order, so after partitioning around pivot `8` you could also end up with something like this:

	[ 3, 0, 5, 2, -1, 1, 8, 8, 14, 26, 10, 27, 9 ]

The only guarantee is that to the left of the pivot are all the smaller elements and to the right are all the larger elements. Because partitioning can change the original order of equal elements, quicksort does not produce a "stable" sort (unlike [merge sort](../merge/), for example). Most of the time that's not a big deal.

## Lomuto's partitioning scheme

Here's an implementation of Lomuto's partitioning scheme:

```go
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
```


```go
numbers := []int { 10, 0, 3, 9, 2, 14, 26, 27, 1, 5, 8, -1, 8 }
p := partitionLomuto(numbers, 0, len(numbers)-1)
```

The `low` and `high` parameters are necessary because when this is used inside quicksort, you don't always want to (re)partition the entire array, only a limited range that becomes smaller and smaller.

Previously we used the middle array element as the pivot but it's important to realize that the Lomuto algorithm always uses the *last* element, `a[high]`, for the pivot. Because we've been pivoting around `8` all this time, I swapped the positions of `8` and `26` in the example so that `8` is at the end of the array and is used as the pivot value here too.

After partitioning, the array looks like this:

	[ 0, 3, 2, 1, 5, 8, -1, 8, 9, 10, 14, 26, 27 ]
	                        *

The variable `p` contains the return value of the call to `partitionLomuto()` and is 7. This is the index of the pivot element in the new array (marked with a star).

The left partition goes from 0 to `p-1` and is `[ 0, 3, 2, 1, 5, 8, -1 ]`. The right partition goes from `p+1` to the end, and is `[ 9, 10, 14, 26, 27 ]` (the fact that the right partition is already sorted is a coincidence).

You may notice something interesting... The value `8` occurs more than once in the array. One of those `8`s did not end up neatly in the middle but somewhere in the left partition. That's a small downside of the Lomuto algorithm as it makes quicksort slower if there are a lot of duplicate elements.

So how does the Lomuto algorithm actually work? The magic happens in the `for` loop. This loop divides the array into four regions:

1. `a[low...i]` contains all values <= pivot
2. `a[i+1...j-1]` contains all values > pivot
3. `a[j...high-1]` are values we haven't looked at yet
4. `a[high]` is the pivot value

In ASCII art the array is divided up like this:

	[ values <= pivot | values > pivot | not looked at yet | pivot ]
	  low           i   i+1        j-1   j          high-1   high

The loop looks at each element from `low` to `high-1` in turn. If the value of the current element is less than or equal to the pivot, it is moved into the first region using a swap.


After the loop is over, the pivot is still the last element in the array. So we swap it with the first element that is greater than the pivot. Now the pivot sits between the <= and > regions and the array is properly partitioned.

Let's step through the example. The array we're starting with is:

	[| 10, 0, 3, 9, 2, 14, 26, 27, 1, 5, 8, -1 | 8 ]
	   low                                       high
	   i
	   j

Initially, the "not looked at" region stretches from index 0 to 11. The pivot is at index 12. The "values <= pivot" and "values > pivot" regions are empty, because we haven't looked at any values yet.

Look at the first value, `10`. Is this smaller than the pivot? No, skip to the next element.	  

	[| 10 | 0, 3, 9, 2, 14, 26, 27, 1, 5, 8, -1 | 8 ]
	   low                                        high
	   i
	       j

Now the "not looked at" region goes from index 1 to 11, the "values > pivot" region contains the number `10`, and "values <= pivot" is still empty.

Look at the second value, `0`. Is this smaller than the pivot? Yes, so swap `10` with `0` and move `i` ahead by one.

	[ 0 | 10 | 3, 9, 2, 14, 26, 27, 1, 5, 8, -1 | 8 ]
	  low                                         high
	      i
	           j

Now "not looked at" goes from index 2 to 11, "values > pivot" still contains `10`, and "values <= pivot" contains the number `0`.

Look at the third value, `3`. This is smaller than the pivot, so swap it with `10` to get:

	[ 0, 3 | 10 | 9, 2, 14, 26, 27, 1, 5, 8, -1 | 8 ]
	  low                                         high
	         i
	             j

The "values <= pivot" region is now `[ 0, 3 ]`. Let's do one more... `9` is greater than the pivot, so simply skip ahead:

	[ 0, 3 | 10, 9 | 2, 14, 26, 27, 1, 5, 8, -1 | 8 ]
	  low                                         high
	         i
	                 j

Now the "values > pivot" region contains `[ 10, 9 ]`. If we keep going this way, then eventually we end up with:

	[ 0, 3, 2, 1, 5, 8, -1 | 27, 9, 10, 14, 26 | 8 ]
	  low                                        high
	                         i                   j

The final thing to do is to put the pivot into place by swapping `a[i]` with `a[high]`:

	[ 0, 3, 2, 1, 5, 8, -1 | 8 | 9, 10, 14, 26, 27 ]
	  low                                       high
	                         i                  j

And we return `i`, the index of the pivot element.

> **Note:** If you're still not entirely clear on how the algorithm works, I suggest you play with this in the playground to see exactly how the loop creates these four regions.

Let's use this partitioning scheme to build quicksort. Here's the code:

```go
func SortLumoto(arr []int, low, high int) {
	if low < high {
		p := partitionLumoto(arr, low, high)
		SortLumoto(arr, low, p-1)
		SortLumoto(arr, p+1, high)
	}
}
```

This is now super simple. We first call `partitionLomuto()` to reorder the array around the pivot (which is always the last element from the array). And then we call `quicksortLomuto()` recursively to sort the left and right partitions.


```go
numbers = []int {10, 0, 3, 9, 2, 14, 26, 27, 1, 5, 8, -1, 8 }
quicksortLomuto(numbers, 0, len(numbers)-1)
```

Lomuto's isn't the only partitioning scheme but it's probably the easiest to understand. It's not as efficient as Hoare's scheme, which requires fewer swap operations.

## Hoare's partitioning scheme

This partitioning scheme is by Hoare, the inventor of quicksort.

Here is the code:

```go
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
```


```go
numbers := []int{ 8, 0, 3, 9, 2, 14, 10, 27, 1, 5, 8, -1, 26 }
p := partitionHoare(numnbers, 0, len(numbers)-1)
```

Note that with Hoare's scheme, the pivot is always expected to be the *first* element in the array, `arr[low]`. Again, we're using `8` as the pivot element.

The result is:

	[ -1, 0, 3, 8, 2, 5, 1, 27, 10, 14, 9, 8, 26 ]

Note that this time the pivot isn't in the middle at all. Unlike with Lomuto's scheme, the return value is not necessarily the index of the pivot element in the new array.

Instead, the array is partitioned into the regions `[low...p]` and `[p+1...high]`. Here, the return value `p` is 6, so the two partitions are `[ -1, 0, 3, 8, 2, 5, 1 ]` and `[ 27, 10, 14, 9, 8, 26 ]`.

The pivot is placed somewhere inside one of the two partitions, but the algorithm doesn't tell you which one or where. If the pivot value occurs more than once, then some instances may appear in the left partition and others may appear in the right partition.

Because of these differences, the implementation of Hoare's quicksort is slightly different:

```go
func SortHoare(arr []int, low, high int) {
	if low < high {
		p := partitionHoare(arr, low, high)
		SortHoare(arr, low, p)
		SortHoare(arr, p+1, high)
	}
}
```

I'll leave it as an exercise for the reader to figure out exactly how Hoare's partitioning scheme works. :-)

## Picking a good pivot

Lomuto's partitioning scheme always chooses the last array element for the pivot. Hoare's scheme uses the first element. But there is no guarantee that these pivots are any good.

Here is what happens when you pick a bad value for the pivot. Let's say the array is,

	[ 7, 6, 5, 4, 3, 2, 1 ]

and we're using Lomuto's scheme. The pivot is the last element, `1`. After pivoting, we have the following arrays:

	   less than pivot: [ ]
	    equal to pivot: [ 1 ]
	greater than pivot: [ 7, 6, 5, 4, 3, 2 ]

Now recursively partition the "greater than" subarray and get:

	   less than pivot: [ ]
	    equal to pivot: [ 2 ]
	greater than pivot: [ 7, 6, 5, 4, 3 ]

And again:

	   less than pivot: [ ]
	    equal to pivot: [ 3 ]
	greater than pivot: [ 7, 6, 5, 4 ]

And so on...

That's no good, because this pretty much reduces quicksort to the much slower insertion sort. For quicksort to be efficient, it needs to split the array into roughly two halves.

The optimal pivot for this example would have been `4`, so we'd get:

	   less than pivot: [ 3, 2, 1 ]
	    equal to pivot: [ 4 ]
	greater than pivot: [ 7, 6, 5 ]

You might think this means we should always choose the middle element rather than the first or the last, but imagine what happens in the following situation:

	[ 7, 6, 5, 1, 4, 3, 2 ]

Now the middle element is `1` and that gives the same lousy results as in the previous example.

Ideally, the pivot is the *median* element of the array that you're partitioning, i.e. the element that sits in the middle of the sorted array. Of course, you won't know what the median is until after you've sorted the array, so this is a bit of a chicken-and-egg problem. However, there are some tricks to choose good, if not ideal, pivots.

One trick is "median-of-three", where you find the median of the first, middle, and last element in this subarray. In theory that often gives a good approximation of the true median.

Another common solution is to choose the pivot randomly. Sometimes this may result in choosing a suboptimal pivot but on average this gives very good results.

Here is how you can do quicksort with a randomly chosen pivot:

```go
func SortRandom(arr []int, low, high int) {
	if low < high {
		pivotIndex := random(low, high)
		arr[pivotIndex], arr[high] = arr[high], arr[pivotIndex]

		p := partitionLumoto(arr, low, high)
		SortRandom(arr, low, p-1)
		SortRandom(arr, p+1, high)
	}
}
```

There are two important differences with before:

1. The `random(low, high)` function returns an integer in the range `min...max`, inclusive. This is our pivot index.

2. Because the Lomuto scheme expects `arr[high]` to be the pivot entry, we swap `arr[pivotIndex]` with `arr[high]` to put the pivot element at the end before calling `partitionLomuto()`.

It may seem strange to use random numbers in something like a sorting function, but it is necessary to make quicksort behave efficiently under all circumstances. With bad pivots, the performance of quicksort can be quite horrible, **O(n^2)**. But if you choose good pivots on average, for example by using a random number generator, the expected running time becomes **O(n log n)**, which is as good as sorting algorithms get.

## Dutch national flag partitioning

But there are more improvements to make! In the first example of quicksort I showed you, we ended up with an array that was partitioned like this:

	[ values < pivot | values equal to pivot | values > pivot ]

But as you've seen with the Lomuto partitioning scheme, if the pivot occurs more than once the duplicates end up in the left half. And with Hoare's scheme the pivot can be all over the place. The solution to this is "Dutch national flag" partitioning.

The code for this scheme is:

```go
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
```

This works very similarly to the Lomuto scheme, except that the loop partitions the array into four (possibly empty) regions:

- `[low ... smaller-1]` contains all values < pivot
- `[smaller ... equal-1]` contains all values == pivot
- `[equal ... larger]` contains all values > pivot
- `[larger ... high]` are values we haven't looked at yet

Note that this doesn't assume the pivot is in `a[high]`. Instead, you have to pass in the index of the element you wish to use as a pivot.

An example of how to use it:

```go
numbers := []int{ 10, 0, 3, 9, 2, 14, 8, 27, 1, 5, 8, -1, 26 }
partitionDutchFlag(numbers, 0, len(numbers)-1, 10)
```

Just for fun, we're giving it the index of the other `8` this time. The result is:

	[ -1, 0, 3, 2, 5, 1, 8, 8, 27, 14, 9, 26, 10 ]

Notice how the two `8`s are in the middle now. The return value from `partitionDutchFlag()` is `(6, 7)`. This is the range that contains the pivot value(s).

Here is how you would use it in quicksort:

```go
func SortDutchFlag(arr []int, low, high int) {
	if low < high {
		pivotIndex := random(low, high)
		p, q := partitionDutchFlag(arr, low, high, pivotIndex)
		SortDutchFlag(arr, low, p-1)
		SortDutchFlag(arr, q+1, high)
	}
}
```

Using Dutch flag partitioning makes for a more efficient quicksort if the array contains many duplicate elements.

## See also

[Quicksort on Wikipedia](https://en.wikipedia.org/wiki/Quicksort)

*Written by [brotherpowers](https://www.brotherpowers.com/)*
