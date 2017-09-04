package quick

func Swap(arr *[]int, i, j int) {
	(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
}

// Lomuto's partitioning scheme
func PartitionLumoto(arr []int, low, high int) int {
	length := len(arr)

	if length == 0 {
		panic("Array size can not be 0")
	}

	if length == 1 {
		return -1
	}

	i := low
	pivot := arr[high]

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			Swap(&arr, i, j)
			i++
		}
	}

	Swap(&arr, i, high)
	return i
}

// Sort using Lomuto's partitioning
func SortLumoto(arr []int, low, high int) {
	if low < high {
		p := PartitionLumoto(arr, low, high)
		SortLumoto(arr, low, p-1)
		SortLumoto(arr, p+1, high)
	}
}

// Hoare's partitioning scheme
func PartitionHoare(arr []int, low, high int) int {
	length := len(arr)

	if length == 0 {
		panic("Array size is 0")
	}

	pivot := arr[low]
	i := low - 1
	j := high + 1

	for {
		for {
			j--

			if arr[j] <= pivot {
				break
			}
		}

		for {
			i++
			if arr[i] >= pivot {
				break
			}
		}

		if i < j {
			Swap(&arr, i, j)
		} else {
			return j
		}
	}
}

func SortHoare(arr []int, low, high int) {
	if low < high {
		p := PartitionHoare(arr, low, high)
		SortHoare(arr, low, p)
		SortHoare(arr, p+1, high)
	}
}
