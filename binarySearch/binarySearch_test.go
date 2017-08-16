package binarySearch

import "testing"

func TestSearch(t *testing.T){
  var cases[1]TestCase

  arr := []Comparable{Rect{100,200}, Rect{200,200}, Rect{300,400}}
  searchable := Rect{100,200}
  want := 0

  cases[0] = TestCase{arr, searchable, want}
	
  for _,c := range cases{
  	got := Search(c.src, c.searchable)

	if got != c.want{
		t.Errorf("Search(%q) in %q, want %q", c.searchable, c.src, c.want)
	}
  }
}

type TestCase struct{
	src []Comparable
	searchable Comparable
	want int
}

type Rect struct{
	width, height int
}

func (rect Rect) Compare(other Comparable) int{
	value, ok := other.(Rect)

	if !ok{
		panic("The other must be of type Rect")
	}

	areaRect := rect.width * rect.height
	areaOther := value.width * value.height

	if areaRect == areaOther{
		return 0
	}else if areaRect > areaOther{
		return 1
	}else{
		return -1
	}
} 

