package shell

func Sort(arr []int) {
	length := len(arr)

	if length == 0 {
		panic("Array size can not be zero")
	}

	if length == 1 {
		return
	}

	for subListCount := length / 2; subListCount > 0; subListCount /= 2 {
		for i := subListCount; i < length; i++ {
			for j := i; j >= subListCount && arr[j-subListCount] > arr[j]; j -= subListCount {
				arr[j], arr[j-subListCount] = arr[j-subListCount], arr[j]
			}
		}
	}
}
