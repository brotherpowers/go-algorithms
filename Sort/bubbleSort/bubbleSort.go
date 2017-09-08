package bubbleSort

// Sort bubbleSort sorting function
func Sort(inputArray []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i <= len(inputArray)-1; i++ {
			if inputArray[i-1] > inputArray[i] {
				inputArray[i], inputArray[i-1] = inputArray[i-1], inputArray[i]
				swapped = true
			}
		}
	}

	return inputArray
}
