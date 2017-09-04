package merge

func Sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	midIndex := len(arr) / 2
	leftArray := Sort(arr[:midIndex])
	rightArray := Sort(arr[midIndex:])

	return merge(leftArray, rightArray)
}

func merge(leftPile, rightPile []int) []int {

	leftIndex := 0
	rightIndex := 0

	var orderedPile []int

	for leftIndex < len(leftPile) && rightIndex < len(rightPile) {
		if leftPile[leftIndex] < rightPile[rightIndex] {
			orderedPile = append(orderedPile, leftPile[leftIndex])
			leftIndex++

		} else if leftPile[leftIndex] > rightPile[rightIndex] {
			orderedPile = append(orderedPile, rightPile[rightIndex])
			rightIndex++
		} else {
			orderedPile = append(orderedPile, leftPile[leftIndex])
			leftIndex++
			orderedPile = append(orderedPile, rightPile[rightIndex])
			rightIndex++
		}

	}

	for leftIndex < len(leftPile) {
		orderedPile = append(orderedPile, leftPile[leftIndex])
		leftIndex++
	}

	for rightIndex < len(rightPile) {
		orderedPile = append(orderedPile, rightPile[rightIndex])
		rightIndex++
	}

	return orderedPile
}
