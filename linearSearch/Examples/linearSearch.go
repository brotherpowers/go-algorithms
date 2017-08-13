package main

import "fmt"
import "github.com/brotherpowers/linearSearch"

func main() {
	fmt.Println("vim-go")

	arr := []interface{}{1,2,3,4}

	position := linearSearch.Search(arr, 4)

	fmt.Println("Search at: ", position)
}

