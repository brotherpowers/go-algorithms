package binarySearch

// The iterative version of binary search

func SearchIterative(inputArray []int, searchTerm int) int {
	lowerBound := 0
	upperBound := len(inputArray)

	for lowerBound < upperBound {
		mid := lowerBound + (upperBound-lowerBound)/2

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

// The recursive version of binary search.

func SearchRecursive(inputArray []int, searchTerm, lowerBound, upperBound int) int {
	if lowerBound > upperBound {
		return -1
	} else {
		mid := lowerBound + (upperBound-lowerBound)/2

		if inputArray[mid] > searchTerm {
			return SearchRecursive(inputArray, searchTerm, lowerBound, mid)
		} else if inputArray[mid] < searchTerm {
			return SearchRecursive(inputArray, searchTerm, mid+1, upperBound)
		} else {
			return mid
		}
	}
	return -1
}
