package main

import "fmt"
import "github.com/brotherpowers/binarySearch"

func main() {
	arr := []binarySearch.Comparable{Rect{100,200}, Rect{200, 200}, Rect{200,350}}

	searchable := Rect{100,200}

	result := binarySearch.Search(arr, searchable)

	fmt.Println(result)
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
