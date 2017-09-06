package radixSort

import (
	"reflect"
	"testing"
)

// Radix Sort
func TestSort(t *testing.T) {
	cases := []struct {
		in   []int
		want []int
	}{
		{[]int{23, 11, 4, 12, 9, 7, 4, 15, 17, 0}, []int{0, 4, 4, 7, 9, 11, 12, 15, 17, 23}},
		{[]int{170, 45, 75, 90, 802, 24, 2, 66}, []int{2, 24, 45, 66, 75, 90, 170, 802}},
	}

	for _, c := range cases {
		Sort(c.in)
		got := c.in
		if !reflect.DeepEqual(c.in, c.want) {
			t.Errorf("Sort -> %v, want = %v", got, c.want)
		}
	}
}
