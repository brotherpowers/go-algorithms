package minmax

import "testing"

func TestMaximum(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{1, 2, 3, 4, 7, 12, 19, 27, 28, 19, 8, 67, 11, 83}, 83},
		{[]int{1, 6, 3, 7, 12, 19, 47, 38, 9, 18, 67, 1, 67}, 67},
		{[]int{-8, -24, 0, 12, 13, 42, 41, 9}, 42},
		{[]int{-4, -7, -12, -19, -27, -28, -19, -8, -67, -11, -83}, -4},
	}

	for _, c := range cases {
		got := Maximum(c.in)

		if got != c.want {
			t.Errorf("Maximum in %v == %v, want %v", c.in, got, c.want)
		}
	}
}
