package main

import (
	"fmt"
	"math"

)


func main() {
	a:= [] int{ 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610 }
	key:= 611

	JumpSearch(a,key)

}
func JumpSearch(inputArray [] int , searchElement int) int{

	length:= len(inputArray)
	steps:= int(math.Sqrt(float64(length)))
	fmt.Println(steps)
	low := 0
	high:=1
	for inputArray[min(high, length)-1] < searchElement {
		low = high
		high = high +steps
		if high >= length {
			fmt.Println("Element is not Found")
			return -1
		}
	}

	res:= BinarySearch(inputArray,low,high,searchElement)
	if res==-1{
		fmt.Println("Element not Found")
		return -1
	}else{
		fmt.Println("Element found at " , res)
		return res
	}

}

func BinarySearch(inputArray [] int , low int , high int ,searchElement int) int {

	if low>high{
		return -1
	}

	mid:= (high+low)/2

	if(inputArray[mid] == searchElement){
		return mid
	}else if(inputArray[mid] > searchElement){
		return BinarySearch(inputArray,low,mid,searchElement)
	}else{
		return BinarySearch(inputArray,mid+1,high,searchElement)
	}

}

func min(first int, second int) int {
	if first < second {
		return first
	}

	return second
}