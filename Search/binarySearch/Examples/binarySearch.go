package main

import "fmt"
import "github.com/go-algorithms/Search/binarySearch"

func main() {
	intArr := []int{1,2,3,17,45,29,400,1032,9979}
	searchNumber := 47

	// log
	fmt.Printf("Searching for (%v) in %v\n", searchNumber, intArr)
	
	searchResult := binarySearch.Search(intArr, searchNumber)
	
	if searchResult == -1{
		fmt.Printf("Element %v not found in %v",searchNumber, intArr)
	}else{
		fmt.Printf("element found at  (%v) is %v\n", searchNumber, searchResult)
	}
}

