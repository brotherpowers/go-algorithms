package interpolationSearch

func Search(inputArray []int, searchTerm int) int {

	length := len(inputArray)
	low := 0
	high := length - 1

	for inputArray[high] != inputArray[low] && searchTerm >= inputArray[low] && searchTerm <= inputArray[high] {

		mid := low + ((searchTerm - inputArray[low]) * (high - low) / (inputArray[high] - inputArray[low]))

		if inputArray[mid] < searchTerm {
			low = mid + 1
		} else if searchTerm < inputArray[mid] {
			high = mid - 1
		}

		return mid
	}

	if searchTerm == inputArray[low] {
		return low
	}

	return -1
}
