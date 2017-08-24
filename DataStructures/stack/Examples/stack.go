package main

import (
	"fmt"

	"github.com/go-algorithms/DataStructures/stack"
)

func main() {
	fmt.Println("vim-go")

	// Array of interface
	numbers := []interface{}{11, 59, 3, 2, 53, 17, 31, 7, 19, 67, 47, 13, 37, 61, 29, 43, 5, 41, 23}

	// Stack
	stack := stack.Stack{numbers}

	// Log the original Stack
	fmt.Println("Original Stack:", stack)

	// Push 22 onto Stack
	stack.Push(22)

	// log After push operation
	fmt.Println("PUSH (1):", stack)
	stack.Pop()
	stack.Pop()
	stack.Pop()

	fmt.Println("Stack", stack)
}
