package jumpSearch

import "math"

// Search with jump
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

// Returns square root of a integer number
func sqrt(number int) int {
	return int(math.Sqrt(float64(number)))
}

// Returns min of two numbers
func min(first int, second int) int {
	if first < second {
		return first
	}

	return second
}
