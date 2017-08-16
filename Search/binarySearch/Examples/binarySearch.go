package main

import "fmt"
import "github.com/brotherpowers/Search/binarySearch"

func main() {
	arr := []binarySearch.Comparable{Rect{100,200}, Rect{200, 200}, Rect{200,350}}
	searchable := Rect{100,200}

	// log
	fmt.Printf("Searching for (%v) in %v\n", searchable, arr)

	result := binarySearch.Search(arr, searchable)

	// log
	fmt.Printf("Element found at %v\n", result)
	
	intArr := []int{1,2,3,17,45,29,400,1032,9979}
	searchNumber_1 := 47

	// log
	fmt.Printf("Searching for (%v) in %v\n", searchNumber_1, intArr)
	searchResultNumber_1 := binarySearch.SearchInt(intArr, searchNumber_1)
	
	// log
	fmt.Printf("element found at  (%v) is %v\n", searchNumber_1, searchResultNumber_1)

}

type Rect struct{
	width, height int
}

func (r Rect) String() string{
	return fmt.Sprintf("Rect{%v,%v}", r.width, r.height)
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
