package gnomeSort

func Sort(arr []int) {
	length := len(arr)
	if length < 1 {
		return
	}

	i := 0
	for i < length {

		if i == 0 || arr[i] >= arr[i-1] {
			i++
		} else {
			arr[i], arr[i-1] = arr[i-1], arr[i]
			i--
		}
	}
}
