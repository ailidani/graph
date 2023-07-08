package stack

// stack implements the stack data structure backed by single linked list
type stack[T any] struct {
	top  *node[T]
	size int
}

type node[T any] struct {
	value T
	prev  *node[T]
}

// New creates a new stack
func New[T any]() *stack[T] {
	return new(stack[T])
}

// Len returns the number of items in the stack
func (s *stack[T]) Size() int {
	return s.size
}

// Peek views the top item on the stack
func (s *stack[T]) Peek() T {
	if s.size == 0 {
		var zero T
		return zero
	}
	return s.top.value
}

// Pop the top item of the stack and return it
func (s *stack[T]) Pop() T {
	if s.size == 0 {
		var zero T
		return zero
	}

	n := s.top
	s.top = n.prev
	s.size--
	return n.value
}

// Push a value onto the top of the stack
func (s *stack[T]) Push(value T) {
	n := &node[T]{value, s.top}
	s.top = n
	s.size++
}

func (s *stack[T]) Empty() bool {
	return s.size == 0
}
