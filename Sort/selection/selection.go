package selection

func Sort(arr []int) {
	length := len(arr)

	if length == 0 {
		panic("Array size can not be zero")
	}

	if length == 1 {
		// array is already sorted
		return
	}

	for i := 0; i < length-1; i++ {
		min := i

		for j := i + 1; j < length; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		if i != min {
			temp := arr[i]
			arr[i] = arr[min]
			arr[min] = temp
		}
	}
}
