package binarySearch

func Search(inputArray []int, searchTerm int) int{
	lowerBound := 0
	upperBound := len(inputArray)

	for lowerBound < upperBound{
		mid := lowerBound + (upperBound - lowerBound) / 2
	
		if inputArray[mid] == searchTerm{
			return mid
		}else if inputArray[mid] < searchTerm{
			lowerBound = mid + 1
		}else{
			upperBound = mid
		}
	}
	return -1
}
