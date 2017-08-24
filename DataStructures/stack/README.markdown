# Stack

A stack is like an array but with limited functionality. You can only *push* to add a new element to the top of the stack, *pop* to remove the element from the top, and *peek* at the top element without popping it off.

Why would you want to do this? Well, in many algorithms you want to add objects to a temporary list at some point and then pull them off this list again at a later time. Often the order in which you add and remove these objects matters.

A stack gives you a LIFO or last-in first-out order. The element you pushed last is the first one to come off with the next pop. (A very similar data structure, the [queue](../Queue/), is FIFO or first-in first-out.)

For example, let's push a number on the stack:

```go
stack.Push(10)
```

The stack is now `[ 10 ]`. Push the next number:

```go
stack.Push(3)
```

The stack is now `[ 10, 3 ]`. Push one more number:

```go
stack.Push(57)
```

The stack is now `[ 10, 3, 57 ]`. Let's pop the top number off the stack:

```go
stack.Pop()
```

This returns `57`, because that was the most recent number we pushed. The stack is `[ 10, 3 ]` again.

```go
stack.Pop()
```

This returns `3`, and so on. If the stack is empty, popping gives an error ("Stack Underflow").

A stack is just a wrapper around an array that just lets you push, pop, and look at the top element of the stack:

```go
import "errors"

type Stack struct {
	Array []interface{}
}

func (s Stack) IsEmpty() bool {
	return len(s.Array) == 0
}

func (s Stack) Count() int {
	return len(s.Array)
}

func (s *Stack) Push(element interface{}) {
	s.Array = append(s.Array, element)
}

func (s *Stack) Pop() (interface{}, error) {
	length := len(s.Array)
	if length == 0 {
		return 0, errors.New("Stack Underflow")
	}
	element := s.Array[length-1]
	s.Array = s.Array[:length-1]
	return element, nil
}
```

Notice that a push puts the new element at the end of the array, not the beginning. Inserting at the beginning of an array is expensive, an **O(n)** operation, because it requires all existing array elements to be shifted in memory. Adding at the end is **O(1)**; it always takes the same amount of time, regardless of the size of the array.


*Written by [brotherpowers](https://www.brotherpowers.com/)*