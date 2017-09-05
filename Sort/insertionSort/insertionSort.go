package insertionSort

func Sort(arr []int) {
	length := len(arr)

	// Array must have at-least one element
	if length == 0 {
		panic("Array size 0")
	}

	// Return if we have nothing to sort
	if length == 1 {
		return
	}

	for i := 1; i < length; i++ {
		j := i
		temp := arr[j]

		for j > 0 && temp < arr[j-1] {
			arr[j] = arr[j-1]
			j--
		}
		arr[j] = temp
	}
}
