package countingSort

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
		{[]int{10, 9, 8, 7, 1, 2, 7, 3}, []int{1, 2, 3, 7, 7, 8, 9, 10}},
	}

	for _, c := range cases {

		got := Sort(c.in) // Sort the array

		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Sort == %v, want = %v", got, c.want)
		}
	}

}
