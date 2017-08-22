package stack

import (
	"testing"
	"reflect"
)

// PUSH
func TestPush(t *testing.T){
	cases := []struct{
	in Stack
	push int
	want Stack
	}{
		{Stack{[]int{1, 2, 46, 74, 89, 105}}, 200, Stack{[]int{1, 2, 46, 74, 89, 105, 200}}},
		{Stack{[]int{53, 77, 88, 65, 42, 31, 33, 33, 20}}, 3, Stack{[]int{53, 77, 88, 65, 42, 31, 33, 33, 20, 3}}},
		{Stack{[]int{1, 2, 46, 74, 89, 105}}, 6, Stack{[]int{1, 2, 46, 74, 89, 105, 6}}},
	}

	for _, c := range cases{

		// PUSH
		c.in.Push(c.push)	

		// result is the stack 
		got := c.in

		if !reflect.DeepEqual(got, c.want){
			t.Errorf("Stack after  PUSH(%v) = %v and want = %v", got, c.push, c.want)	
		}
	}
}

// POP

func TestPop(t *testing.T){
	cases := []struct{
	in Stack
	repeat int
	want Stack
	}{
		{Stack{[]int{1, 2, 46, 74, 89, 105}}, 1, Stack{[]int{1, 2, 46, 74, 89}}},
		{Stack{[]int{53, 77, 88, 65, 42, 31, 33, 33, 20}}, 3, Stack{[]int{53, 77, 88, 65, 42, 31}}},
		{Stack{[]int{1, 2, 46, 74, 89, 105}}, 6, Stack{[]int{}}},
	}

	for _, c := range cases{

		// POP
		for i := 0; i < c.repeat; i++{ 
			c.in.Pop()
		}

		// result is the stack 
		got := c.in

		if !reflect.DeepEqual(got, c.want){
			t.Errorf("Stack after POP %v times = %v and want = %v", got, c.repeat, c.want)	
		}
	}
}
