package queue

type Queue[T any] interface {
	Push(T)
	Pop() T
	Front() T
	Back() T
	//Contains(T) bool
	//Remove(T)
	Size() int
}

func New[T any]() Queue[T] {
	return &queue[T]{
		queue: make([]T, 128),
		size:  0,
		head:  0,
		tail:  0,
	}
}

type queue[T any] struct {
	queue []T
	size  int
	head  int
	tail  int
}

func (q *queue[T]) Size() int {
	return q.size
}

func (q *queue[T]) Push(e T) {
	if q.head == q.tail && q.size > 0 {
		queue := make([]T, len(q.queue)*2)
		copy(queue, q.queue[q.head:])
		copy(queue[len(q.queue)-q.head:], q.queue[:q.head])
		q.head = 0
		q.tail = len(q.queue)
		q.queue = queue
	}
	q.queue[q.tail] = e
	q.tail = (q.tail + 1) % len(q.queue)
	q.size++
}

func (q *queue[T]) Pop() T {
	if q.size == 0 {
		var zero T
		return zero
	}
	e := q.queue[q.head]
	q.head = (q.head + 1) % len(q.queue)
	q.size--
	return e
}

func (q queue[T]) Front() T {
	if q.size == 0 {
		var zero T
		return zero
	}
	return q.queue[q.head]
}

func (q queue[T]) Back() T {
	if q.size == 0 {
		var zero T
		return zero
	}
	return q.queue[q.tail-1]
}

/*
func (q queue[T]) Contains(e T) bool {
	if q.size == 0 {
		return false
	}

	for _, x := range q.queue {
		if e == x {
			return true
		}
	}

	return false
}

func (q *queue[T]) Remove(e T) {
	if q.size == 0 {
		return
	}

	for i, x := range q.queue {
		if x == e {
			q.queue = append(q.queue[:i], q.queue[i+1:]...)
		}
	}
}
*/
