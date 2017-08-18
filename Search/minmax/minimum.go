package minmax

func Minimum(inputArray []int) int {
	if len(inputArray) == 0 {
		panic("Array must have at-least one element")
	}

	min := inputArray[0]

	for _, value := range inputArray {
		if value < min {
			min = value
		}
	}

	return min
}
