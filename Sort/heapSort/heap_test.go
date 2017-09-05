package heapSort

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {

	cases := []struct {
		in   []int
		want []int
	}{
		{[]int{23, 11, 4, 12, 9, 7, 4, 15, 17, 0}, []int{0, 4, 4, 7, 9, 11, 12, 15, 17, 23}},
		{[]int{10, -1, 3, 9, 2, 27, 8, 5, 1, 3, 0, 26}, []int{-1, 0, 1, 2, 3, 3, 5, 8, 9, 10, 26, 27}},
	}

	for _, c := range cases {

		Sort(c.in) // Sort the array

		got := c.in // Fetch the Sorted array

		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Sort == %v, want = %v", got, c.want)
		}
	}

}
