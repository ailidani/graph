package stack

import (
	"testing"
)

func Test(t *testing.T) {
	s := New[int]()

	if s.Size() != 0 {
		t.Errorf("Length of an empty stack should be 0")
	}

	s.Push(1)

	if s.Size() != 1 {
		t.Errorf("Length should be 0")
	}

	if s.Peek() != 1 {
		t.Errorf("Top item on the stack should be 1")
	}

	if s.Pop() != 1 {
		t.Errorf("Top item should have been 1")
	}

	if s.Size() != 0 {
		t.Errorf("Stack should be empty")
	}

	s.Push(1)
	s.Push(2)

	if s.Size() != 2 {
		t.Errorf("Length should be 2")
	}

	if s.Peek() != 2 {
		t.Errorf("Top of the stack should be 2")
	}
}
