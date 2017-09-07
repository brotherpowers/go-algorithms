# Linear Search

Goal: Find a particular value in an array.

Like [Binary Search](../Search/binarySearch/), Jump Search is a searching algorithm for sorted arrays. The basic idea is to check fewer elements than [Linear search](../Search/linearSearch/) by jumping ahead by fixed steps or skipping some elements in place of searching all elements.

## An example

Let's say we have an array of numbers `[ 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610] ` and we want to check if the array contains the number `55`.

`steps := sqrt(len(numnbers))	// 4`

Initially loop starts at index `0`. It looks for an index an index that holds value greater then what we are searching for. And in every iteration it skips `steps` number of indices.

	[ 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610 ]
	  *

Our head is at index `0` and the value is `0` which is lesser than `55`. So we jump `4` steps.

	[ 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610 ]
	  |---------> *
	  
Our head is now at index `4` and the value is `3` which is lesser than `55`. So we again jump `4` steps.
	  
  	[ 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610 ]
  	  |           |-----------> *
	  
Our head is now at index `8` and the value is `21` which is lesser than `55`. So we again jump `4` steps.
	  
  	[ 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610 ]
  	  |           |            |------------------------------> *

Our head is now at index `15` which is the last position in our array and the value is `610` which is greater than `55`. So we break our loop and we will now perform `linearSearch()` on the optimal region that starts from last feasible index position i.e `8`.

From here on we will perform [Linear Search](../Search/linearSearch/) traversing the array elements one by one till we find an index which holds the value equal to our search or we return -1 if element is not found.

  	[ X, X, X, X, X, X, X, X, 21, 34, 55, 89, 144, 233, 377, 610 ]
  	                         |---|---> X
			  

###What is the optimal block size to be skipped?

In the worst case, we have to do `n/steps` jumps *(where `n` is the size of array and `steps` is the size of the jump we will be using to skip of the elements)* and if the last checked value is greater than the element to be searched for, we perform `steps-1` comparisons more for linear search. Therefore the total number of comparisons in the worst case will be ((n/steps) + steps-1). The value of the function `((n/steps) + steps-1)` will be minimum when `steps = √n`. Therefore, the best step size is steps = √n.

## The code

Here is a simple implementation of jump search:

```go
func Search(inputArray []int, searchElement int) int {
	key := 0
	length := len(inputArray)
	steps := sqrt(length)

	for inputArray[min(steps, length)-1] < searchElement {
		key = steps
		steps = steps + sqrt(length)
		if key >= length {
			return -1
		}
	}

	for inputArray[key] < searchElement {
		key = key + 1

		if key == min(steps, length) {
			return -1
		}
	}

	if inputArray[key] == searchElement {
		return key
	}

	return -1
}
```

## Performance

The optimal size of a block to be jumped is **O(√n)**. This makes the time complexity of Jump Search **O(√ n)**. The time complexity of Jump Search is between [Linear search](../Search/linearSearch/) **O(n)** and [Binary Search](../Search/binarySearch/) **O(Log n)**.

Binary Search is better than Jump Search, but Jump search has an advantage that we traverse back only once. [Binary Search](../Search/binarySearch/) may require up to **O(Log n)** jumps, consider a situation where the element to be search is the smallest element or smaller than the smallest. So in a systems where jumping back is costly, we use Jump Search.


## See also

[Jump search on Wikipedia](https://en.wikipedia.org/wiki/Jump_search)

*Written by [brotherpowers](https://www.brotherpowers.com/)*
