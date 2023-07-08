package list

import (
	"reflect"
	"testing"
)

func TestList(t *testing.T) {
	l := new(List[int])
	l.PushBack(3)
	l.PushBack(2)
	l.PushBack(1)
	l.PushBack(0)

	// 3 -> 2 -> 1 -> 0
	if l.Front().Value != 3 {
		t.Error(l.Front().Value)
	}

	if l.Back().Value != 0 {
		t.Error(l.Back().Value)
	}

	if l.Len() != 4 {
		t.Error(l.Len())
	}

	// 3 -> 2 -> 1
	l.Remove(l.Back())
	if l.Back().Value != 1 {
		t.Error(l.Back().Value)
	}

	if l.Len() != 3 {
		t.Error(l.Len())
	}

	// 3 -> 2 -> 1 -> 4
	l.InsertAfter(4, l.Back())
	if l.Back().Value != 4 {
		t.Error(l.Back().Value)
	}
	if l.Len() != 4 {
		t.Error(l.Len())
	}

	// 1 -> 2 -> 3 -> 4
	l.Sort(func(a, b int) bool {
		return a < b
	})

	array := l.Front().ToArray()
	if !reflect.DeepEqual(array, []int{1, 2, 3, 4}) {
		t.Errorf("%v", array)
	}
}
