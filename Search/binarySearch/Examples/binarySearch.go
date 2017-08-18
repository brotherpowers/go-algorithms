package main

import "fmt"
import "github.com/go-algorithms/Search/binarySearch"

func main() {
	numbers := []int{ 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67}
	searchNumber := 43

	// log
	fmt.Printf("Searching for (%v) in %v\n", searchNumber, numbers)
	
	searchResult := binarySearch.SearchRecursive(numbers, searchNumber, 0, len(numbers)-1)
	
	if searchResult == -1{
		fmt.Printf("Element %v not found in %v\n",searchNumber, numbers)
	}else{
		fmt.Printf("element found at (%v)\n", searchResult)
	}
}

