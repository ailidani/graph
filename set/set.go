package set

import (
	"math/rand"
)

type Set[T any] interface {
	Add(...T)
	Remove(T)
	Clear()
	Contains(T) bool
	GetRandom() T
	Array() []T
	Clone() Set[T]
	Size() int
	Select(func(T) bool) []T
}

type set[T comparable] map[T]struct{}

func New[T comparable]() Set[T] {
	return set[T](make(map[T]struct{}))
}

func (s set[T]) Add(a ...T) {
	for _, e := range a {
		s[e] = struct{}{}
	}
}

func (s set[T]) Merge(a Set[T]) {
	for _, e := range a.Array() {
		s.Add(e)
	}
}

func (s set[T]) Remove(e T) {
	delete(s, e)
}

func (s set[T]) Contains(e T) bool {
	_, exists := s[e]
	return exists
}

// GetRandom returns random element from set
func (s set[T]) GetRandom() T {
	n := rand.Intn(len(s))
	for k := range s {
		n--
		if n == 0 {
			return k
		}
	}
	panic("set get random failed")
}

func (s set[T]) Array() []T {
	a := make([]T, len(s))
	i := 0
	for e := range s {
		a[i] = e
		i++
	}
	return a
}

func (s set[T]) Clear() {
	for k := range s {
		delete(s, k)
	}
}

func (s set[T]) Clone() Set[T] {
	clone := New[T]()
	for v := range s {
		clone.Add(v)
	}
	return clone
}

func (s set[T]) Size() int {
	return len(s)
}

func (s set[T]) Select(f func(T) bool) []T {
	a := make([]T, 0)
	for e := range s {
		if f(e) {
			a = append(a, e)
		}
	}
	return a
}
