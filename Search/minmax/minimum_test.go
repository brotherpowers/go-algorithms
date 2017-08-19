package minmax

import "testing"

func TestMinimum(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{1, 2, 3, 4, 7, 12, 19, 27, 28, 19, 8, 67, 11, 83}, 1},
		{[]int{1, 6, 3, 7, 12, 19, 47, 38, 9, 18, 67, 1, 67}, 1},
		{[]int{-8, -24, 0, 12, 13, 42, 41, 9}, -24},
		{[]int{-4, -7, -12, -19, -27, -28, -19, -8, -67, -11, -83}, -83},
	}

	for _, c := range cases {
		got := Minimum(c.in)

		if got != c.want {
			t.Errorf("Minimum in %v == %v, want %v", c.in, got, c.want)
		}
	}
}
