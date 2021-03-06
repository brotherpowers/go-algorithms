package kthLargest

import (
	"testing"
)

func TestKthLargest(t *testing.T) {
	cases := []struct {
		in           []int
		search, want int
	}{
		{[]int{11, 59, 3, 2, 53, 17, 31, 7, 19, 67, 47, 13, 37, 61, 29, 43, 5, 41, 23}, 2, 59},
		{[]int{5, 2, 4, 7}, 1, 5},
		{[]int{0, 1, 1, 3, 3, 3, 3, 6, 8, 10, 11, 11}, 3, 8},
	}

	for _, c := range cases {
		got := KthLargest(c.in, c.search)

		if got != c.want {
			t.Errorf("%vth largest number in %v == %v, want = %v", c.search, c.in, got, c.want)
		}
	}
}

// Test KthLargestUsingHeap
func TestKthLargestUsingHeap(t *testing.T) {
	cases := []struct {
		in           []int
		search, want int
	}{
		{[]int{11, 59, 3, 2, 53, 17, 31, 7, 19, 67, 47, 13, 37, 61, 29, 43, 5, 41, 23}, 2, 59},
		{[]int{5, 2, 4, 7}, 1, 5},
		{[]int{0, 1, 1, 3, 3, 3, 3, 6, 8, 10, 11, 11}, 3, 8},
	}

	for _, c := range cases {
		got := KthLargestUsingHeap(c.in, c.search)

		if got != c.want {
			t.Errorf("%vth largest number in %v == %v, want = %v", c.search, c.in, got, c.want)
		}
	}
}
