package interpolationSearch

import "testing"

func TestSearch(t *testing.T) {
	cases := []struct {
		in           []int
		search, want int
	}{
		{[]int{1, 2, 46, 74, 89, 105}, 89, 4},
		{[]int{5, 6, 78, 343, 568, 999}, 1000, -1},
		{[]int{0, 0, 23, 34, 450, 550, 555, 679, 843}, 843, 8},
		{[]int{10, 16, 80, 143, 268, 599, 768}, 10, 0},
		{[]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67}, 7, 3},
	}

	for _, c := range cases {
		got := Search(c.in, c.search)

		if got != c.want {
			t.Errorf("Search (%v) in (%v)\n Result = %v,\n Wants = %v", c.search, c.in, got, c.want)
		}
	}
}
