package binarySearch

import "testing"

func TestSearch(t *testing.T){

	cases := []struct{
		in []Comparable
		search Comparable
		want int
	}{
		{[]Comparable{Rect{100,200}, Rect{200,200}, Rect{300,400}}, Rect{100,200}, 0},
		{[]Comparable{Rect{23,92}, Rect{56,20}, Rect{0,40},Rect{999,89}}, Rect{999,89}, 3},
		{[]Comparable{Rect{100,200}, Rect{200,200}, Rect{300,400}}, Rect{10,20}, -1},
		{[]Comparable{Rect{100,200}, Rect{200,200}, Rect{200,300},Rect{300,500}}, Rect{200,200}, 1},
	}
	
  for _,c := range cases{
  	got := Search(c.in, c.search)

	if got != c.want{
		t.Errorf("Search(%v) == %v, want %v", c.in, got, c.want)
	}
  }
}

func TestSearchInt(t *testing.T){
	cases := []struct {
				in []int
				search, want int 
			}{
				{[]int{1,2,46,74,89,105}, 89, 4},
				{[]int{5,6,78,343,568,999}, 1000, -1 },
				{[]int{0,0,23,34,450,550,555,679,843}, 843, 8 },
				{[]int{10,16,80,143,268,599,768}, 10, 0},
		}

		for _, c := range cases{
			got := SearchInt(c.in, c.search)

			if got != c.want{
				t.Errorf("SearchInt(%v) == %v, want %v", c.in, got , c.want)
			}
		}
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

