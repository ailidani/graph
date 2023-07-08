// Package list provides an implementation of a doubly-linked list with a front
// and back. The individual nodes of the list are publicly exposed so that the
// user can have fine-grained control over the list.
package list

// EqualsFn is a function that returns whether 'a' and 'b' are equal.
type EqualsFn[T any] func(a, b T) bool

// LessFn is a function that returns whether 'a' is less than 'b'.
type LessFn[T any] func(a, b T) bool

// List implements a doubly-linked list
type List[V any] struct {
	front, back *Node[V]
	len         int
}

func (l List[V]) Front() *Node[V] {
	return l.front
}

func (l List[V]) Back() *Node[V] {
	return l.back
}

func (l List[V]) Len() int {
	return l.len
}

// PushBack adds 'v' to the end of the list
func (l *List[V]) PushBack(v V) {
	l.PushBackNode(&Node[V]{
		Value: v,
	})
}

// PushFront adds 'v' to the beginning of the list
func (l *List[V]) PushFront(v V) {
	l.PushFrontNode(&Node[V]{
		Value: v,
	})
}

// PushBackNode adds the node 'n' to the back of the list
func (l *List[V]) PushBackNode(n *Node[V]) {
	n.list = l
	n.next = nil
	n.prev = l.back
	if l.back != nil {
		l.back.next = n
	} else {
		l.front = n
	}
	l.back = n
	l.len++
}

// PushFrontNode adds the node 'n' to the front of the list
func (l *List[V]) PushFrontNode(n *Node[V]) {
	n.list = l
	n.next = l.front
	n.prev = nil
	if l.front != nil {
		l.front.prev = n
	} else {
		l.back = n
	}
	l.front = n
	l.len++
}

// Remove removes the node 'n' from the list
func (l *List[V]) Remove(n *Node[V]) V {
	// check if node belongs to list
	if n.list != l {
		return n.Value
	}

	if n.next != nil {
		n.next.prev = n.prev
	} else {
		l.back = n.prev
	}
	if n.prev != nil {
		n.prev.next = n.next
	} else {
		l.front = n.next
	}
	n.next = nil // avoid memory leaks
	n.prev = nil // avoid memory leaks
	n.list = nil
	l.len--
	return n.Value
}

func (l *List[V]) InsertBefore(v V, mark *Node[V]) *Node[V] {
	if mark.list != l {
		return nil
	}
	n := &Node[V]{
		Value: v,
		prev:  mark.prev,
		next:  mark,
		list:  l,
	}
	mark.prev = n
	if n.prev != nil {
		n.prev.next = n
	} else {
		l.front = n
	}
	l.len++
	return n
}

func (l *List[V]) InsertAfter(v V, mark *Node[V]) *Node[V] {
	if mark.list != l {
		return nil
	}
	n := &Node[V]{
		Value: v,
		prev:  mark,
		next:  mark.next,
		list:  l,
	}
	mark.next = n
	if n.next != nil {
		n.next.prev = n
	} else {
		l.back = n
	}
	l.len++
	return n
}

func (l *List[V]) Sort(less LessFn[V]) {
	current := l.Front()
	for current != nil {
		index := current.Next()
		for index != nil {
			if less(index.Value, current.Value) {
				temp := current.Value
				current.Value = index.Value
				index.Value = temp
			}
			index = index.Next()
		}
		current = current.Next()
	}
}

func (l *List[V]) Contains(v V, equal EqualsFn[V]) bool {
	for i := l.Front(); i != nil; i = i.Next() {
		if equal(i.Value, v) {
			return true
		}
	}
	return false
}
