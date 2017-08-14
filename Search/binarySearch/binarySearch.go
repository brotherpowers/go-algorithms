package binarySearch

// Comparable interface
type Comparable interface {
	// -1 Lesser 1 Greater  0 Equals
	Compare(Comparable) int
}

func Search(src []Comparable, searchable Comparable) int {
	lowerBound := 0
	upperBound := len(src)

	for lowerBound < upperBound {
		mid := lowerBound + (upperBound-lowerBound)/2

		compare := src[mid].Compare(searchable)

		if compare == 0 {
			return mid
		} else if compare < 0 {
			lowerBound = mid + 1
		} else {
			upperBound = mid
		}
	}
	return -1
}
