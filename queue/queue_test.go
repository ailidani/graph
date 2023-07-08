package queue

import (
	"testing"
)

func Test(t *testing.T) {
	q := New[int]()

	if q.Size() != 0 {
		t.Error("Size of an empty queue should be 0")
	}

	N := 10

	for i := 0; i < N; i++ {
		q.Push(i)
	}

	if q.Size() != N {
		t.Errorf("Queue size expected %d actual %d", N, q.Size())
	}

	for i := 0; i < N; i++ {
		v := q.Pop()
		if v != i {
			t.Errorf("Queue pop expected %d actual %d", i, v)
		}

		if q.Size() != N-i-1 {
			t.Errorf("Queue size expected %d actual %d", N-i-1, q.Size())
		}
	}
}
