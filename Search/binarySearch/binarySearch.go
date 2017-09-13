package binarySearch

// SearchIterative : The iterative version of binary search
func SearchIterative(inputArray []int, searchTerm int) int {
	lowerBound := 0
	upperBound := len(inputArray)

	for lowerBound < upperBound {
		mid := (upperBound + lowerBound) >> 1

		if inputArray[mid] == searchTerm {
			return mid
		} else if inputArray[mid] < searchTerm {
			lowerBound = mid + 1
		} else {
			upperBound = mid
		}
	}

	return -1
}

// SearchRecursive : The recursive version of binary search.
func SearchRecursive(inputArray []int, searchTerm, lowerBound, upperBound int) int {

	if lowerBound > upperBound {
		return -1
	}

	mid := (upperBound + lowerBound) >> 1

	if inputArray[mid] > searchTerm {
		return SearchRecursive(inputArray, searchTerm, lowerBound, mid)
	} else if inputArray[mid] < searchTerm {
		return SearchRecursive(inputArray, searchTerm, mid+1, upperBound)
	}

	return mid
}
