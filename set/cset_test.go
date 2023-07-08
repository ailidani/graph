package set

import (
	"sync"
	"testing"
)

func TestCSet(t *testing.T) {
	s := NewCSet[int]()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			s.Add(n)
			wg.Done()
		}(i)
	}
	wg.Wait()

	if !s.Contains(5) {
		t.Error("missing element")
	}

	s.Remove(5)
	if s.Contains(5) {
		t.Error("cannot remove element")
	}
}
