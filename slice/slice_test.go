package slice

import (
	"testing"
)

func TestInclude(t *testing.T) {
	list := []string{}
	if Include(list, "a") {
		t.Fail()
	}

	list = []string{"a", "b", "c"}
	if !Include(list, "a") {
		t.Fail()
	}
}

func TestRemove(t *testing.T) {
	list := []string{"a", "b", "c"}
	Remove(list, "a")
	if Include(list, "a") {
		t.Fail()
	}
}

func TestContains(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6}
	if !Contains(list, 2, func(a, b int) bool {
		return a == b
	}) {
		t.Fail()
	}
}

func TestEqual(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{2, 4, 6, 8, 10}

	if Equal(a, b, func(a, b int) bool {
		return a == b
	}) {
		t.Fail()
	}

	if !Equal(a, b, func(a, b int) bool {
		return a*2 == b
	}) {
		t.Fail()
	}
}
