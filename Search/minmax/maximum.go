package minmax

func Maximum(inputArray []int) int {
	if len(inputArray) == 0 {
		panic("Array must have at-least one element")
	}

	max := inputArray[0]

	for _, value := range inputArray {
		if value > max {
			max = value
		}
	}

	return max
}
