package set

import (
	"math/rand"
)

type MultiSet[T comparable] map[T]int

func NewMultiSet[T comparable]() MultiSet[T] {
	return make(MultiSet[T])
}

func (s MultiSet[T]) Add(a ...T) {
	for _, e := range a {
		s[e]++
	}
}

func (s MultiSet[T]) Remove(e T) {
	if s[e] > 1 {
		s[e]--
	} else {
		delete(s, e)
	}
}

func (s MultiSet[T]) Contains(e T) bool {
	_, exists := s[e]
	return exists
}

func (s MultiSet[T]) GetRandom() T {
	n := rand.Intn(s.Size())
	for k := range s {
		n -= s[k]
		if n <= 0 {
			return k
		}
	}
	panic("set get random failed")
}

func (s MultiSet[T]) Array() []T {
	ret := make([]T, 0, s.Size())
	for k, c := range s {
		for i := 0; i < c; i++ {
			ret = append(ret, k)
		}
	}
	return ret
}

func (s MultiSet[T]) Clone() Set[T] {
	c := make(MultiSet[T], len(s))
	for k, v := range s {
		c[k] = v
	}
	return c
}

func (s MultiSet[T]) Size() int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

func (s MultiSet[T]) Select(f func(T) bool) []T {
	ret := make([]T, 0)
	for k, c := range s {
		if f(k) {
			for i := 0; i < c; i++ {
				ret = append(ret, k)
			}
		}
	}
	return ret
}

func (s MultiSet[T]) Clear() {
	for k := range s {
		delete(s, k)
	}
}
