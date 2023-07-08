package heap

// Heap implements a binary heap
type Heap[T any] struct {
	data []T
	less func(T, T) bool
}

func New[T any](less func(T, T) bool) *Heap[T] {
	return &Heap[T]{
		data: make([]T, 0),
		less: less,
	}
}

func From[T any](less func(T, T) bool, t ...T) *Heap[T] {
	return FromSlice(less, t)
}

func FromSlice[T any](less func(T, T) bool, data []T) *Heap[T] {
	n := len(data)
	for i := n/2 - 1; i >= 0; i-- {
		down(data, i, less)
	}

	return &Heap[T]{
		data: data,
		less: less,
	}
}

func (h *Heap[T]) Push(x T) {
	h.data = append(h.data, x)
	up(h.data, len(h.data)-1, h.less)
}

func (h Heap[T]) Size() int {
	return len(h.data)
}

func (h *Heap[T]) Pop() (T, bool) {
	var x T
	if h.Size() == 0 {
		return x, false
	}

	x = h.data[0]

	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	down(h.data, 0, h.less)

	return x, true
}

func (h *Heap[T]) Peek() (T, bool) {
	if h.Size() == 0 {
		var x T
		return x, false
	}

	return h.data[0], true
}

func down[T any](h []T, i int, less func(T, T) bool) {
	for {
		left, right := 2*i+1, 2*i+2
		if left >= len(h) || left < 0 { // `left < 0` in case of overflow
			break
		}

		// find the smallest child
		j := left
		if right < len(h) && less(h[right], h[left]) {
			j = right
		}

		if !less(h[j], h[i]) {
			break
		}

		h[i], h[j] = h[j], h[i]
		i = j
	}
}

func up[T any](h []T, i int, less func(T, T) bool) {
	for {
		parent := (i - 1) / 2
		if i == 0 || !less(h[i], h[parent]) {
			break
		}

		h[i], h[parent] = h[parent], h[i]
		i = parent
	}
}
