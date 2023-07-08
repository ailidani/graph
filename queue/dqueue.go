package queue

import (
	"net/rpc"
	"sync"
)

func NewDistributedQueue[T any](name string) Queue[T] {
	q := &dqueue[T]{
		name: name,
		queue: queue[T]{
			queue: make([]T, 100),
			size:  0,
			head:  0,
			tail:  0,
		},
	}
	err := rpc.RegisterName(name, q)
	if err != nil {
		return nil
	}
	rpc.HandleHTTP()
	return q
}

type dqueue[T any] struct {
	lock sync.RWMutex
	name string
	queue[T]
}

func (q *dqueue[T]) Size() int {
	q.lock.RLock()
	defer q.lock.RUnlock()
	return q.size
}

func (q *dqueue[T]) Push(e T) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.queue.Push(e)
}

func (q *dqueue[T]) Pop() T {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue.Pop()
}
