package list

// Node is a node in the linked list.
type Node[V any] struct {
	Value      V
	prev, next *Node[V]
	list       *List[V]
}

func (n *Node[V]) Next() *Node[V] {
	return n.next
}

func (n *Node[V]) Prev() *Node[V] {
	return n.prev
}

// Each calls 'fn' on every element from this node onward in the list.
func (n *Node[V]) Each(fn func(val V)) {
	for node := n; node != nil; node = node.Next() {
		fn(node.Value)
	}
}

// EachReverse calls 'fn' on every element from this node backward in the list.
func (n *Node[V]) EachReverse(fn func(val V)) {
	for node := n; node != nil; node = node.Prev() {
		fn(node.Value)
	}
}

func (n *Node[V]) ToArray() []V {
	array := make([]V, 0)
	for node := n; node != nil; node = node.Next() {
		array = append(array, node.Value)
	}
	return array
}
