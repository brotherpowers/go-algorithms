package main

import "fmt"
import "github.com/brotherpowers/Search/binarySearch"

func main() {
	arr := []binarySearch.Comparable{Rect{100,200}, Rect{200, 200}, Rect{200,350}}

	searchable := Rect{100,200}

	result := binarySearch.Search(arr, searchable)

	fmt.Println(result)
	
	intArr := []int{1,2,3,17,45,29,400,1032,9979}
	searchNumber_1 := 47
	searchNumber_2 := 1032

	searchResultNumber_1 := binarySearch.SearchInt(intArr, searchNumber_1)
	fmt.Println("Search result for (%q) is %q", searchNumber_1, searchResultNumber_1)

	searchResultNumber_2 := binarySearch.SearchInt(intArr, searchNumber_2)
	fmt.Println("Search result for (%q) is %q", searchNumber_2, searchResultNumber_2)

}

type Rect struct{
	width, height int
}

func (rect Rect) Compare(other binarySearch.Comparable) int{
	 value, ok := other.(Rect)

	if !ok{
		panic("The other must of type Rect")
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
