package countOccurrences

func Count(inputArray []int, searchTerm int) int {

	lowerBound := func() int {
		low := 0
		high := len(inputArray)

		for low < high {
			mid := low + (high-low)/2

			if inputArray[mid] < searchTerm {
				low = mid + 1
			} else {
				high = mid
			}
		}

		return low
	}

	upperBound := func() int {
		low := 0
		high := len(inputArray)

		for low < high {
			mid := low + (high-low)/2

			if inputArray[mid] > searchTerm {
				high = mid
			} else {
				low = mid + 1
			}
		}

		return high
	}

	return upperBound() - lowerBound()
}
