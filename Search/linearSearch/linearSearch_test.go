package linearSearch

import "testing"

func TestSearch(t *testing.T) {
	var cases [1]TestCase

	arr := []interface{}{1, 2, 3, 4}
	searchable := 4
	want := 3
	cases[0] = TestCase{arr, searchable, want}

	for _, c := range cases {
		got := Search(c.src, c.searchable)

		if got != c.want {
			t.Errorf("Search(%q) in %q, want %q", c.searchable, c.src, c.want)
		}
	}
}

type TestCase struct {
	src        []interface{}
	searchable interface{}
	want       int
}
