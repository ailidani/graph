package set

import "sync"

// cset is concurrent set with generic data as interface{}
type cset[T comparable] struct {
	data Set[T]
	sync.RWMutex
}

func NewCSet[T comparable]() Set[T] {
	return &cset[T]{
		data: New[T](),
	}
}

func (s *cset[T]) Add(e ...T) {
	s.Lock()
	defer s.Unlock()
	s.data.Add(e...)
}

func (s *cset[T]) Remove(e T) {
	s.Lock()
	defer s.Unlock()
	s.data.Remove(e)
}

func (s *cset[T]) Contains(e T) bool {
	s.RLock()
	defer s.RUnlock()
	return s.data.Contains(e)
}

func (s *cset[T]) GetRandom() T {
	s.RLock()
	defer s.RUnlock()
	return s.data.GetRandom()
}

func (s *cset[T]) Size() int {
	s.RLock()
	defer s.RUnlock()
	return s.data.Size()
}

func (s *cset[T]) Array() []T {
	s.RLock()
	defer s.RUnlock()
	return s.data.Array()
}

func (s *cset[T]) Clear() {
	s.Lock()
	defer s.Unlock()
	s.data.Clear()
}

func (s *cset[T]) Clone() Set[T] {
	s.RLock()
	defer s.RUnlock()
	return s.data.Clone()
}

func (s *cset[T]) Select(f func(T) bool) []T {
	s.RLock()
	defer s.RUnlock()
	return s.data.Select(f)
}
