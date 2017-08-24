package stack

import (
	"reflect"
	"testing"
)

// PUSH
func TestPush(t *testing.T) {
	cases := []struct {
		in   Stack
		push interface{}
		want Stack
	}{
		{Stack{[]interface{}{1, 2, 46, 74, 89, 105}}, 200, Stack{[]interface{}{1, 2, 46, 74, 89, 105, 200}}},
		{Stack{[]interface{}{53, 77, 88, 65, 42, 31, 33, 33, 20}}, 3, Stack{[]interface{}{53, 77, 88, 65, 42, 31, 33, 33, 20, 3}}},
		{Stack{[]interface{}{1, 2, 46, 74, 89, 105}}, 6, Stack{[]interface{}{1, 2, 46, 74, 89, 105, 6}}},
	}

	for _, c := range cases {

		// PUSH
		c.in.Push(c.push)

		// result is the stack
		got := c.in

		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Stack after  PUSH(%v) = %v and want = %v", got, c.push, c.want)
		}
	}
}

// POP

func TestPop(t *testing.T) {
	cases := []struct {
		in     Stack
		repeat int
		want   Stack
	}{
		{Stack{[]interface{}{1, 2, 46, 74, 89, 105}}, 1, Stack{[]interface{}{1, 2, 46, 74, 89}}},
		{Stack{[]interface{}{53, 77, 88, 65, 42, 31, 33, 33, 20}}, 3, Stack{[]interface{}{53, 77, 88, 65, 42, 31}}},
		{Stack{[]interface{}{1, 2, 46, 74, 89, 105}}, 6, Stack{[]interface{}{}}},
	}

	for _, c := range cases {

		// POP
		for i := 0; i < c.repeat; i++ {
			c.in.Pop()
		}

		// result is the stack
		got := c.in

		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Stack after POP %v times = %v and want = %v", got, c.repeat, c.want)
		}
	}
}
