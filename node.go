package graph

type Node[K comparable] interface {
	ID() K
}

func NewNode[K comparable](id K) Node[K] {
	return &node[K]{
		id: id,
	}
}

//----------
// node
//----------

type node[K comparable] struct {
	id K
}

func (n node[K]) ID() K { return n.id }

//----------
// nodes
//----------

type Iterator interface {
	// Next advances the iterator and returns whether there are more items.
	Next() bool
	// Len returns the number of items remaining in the iterator.
	Len() int
	// Reset returns the iterator to its start position.
	Reset()
}

type Nodes[K comparable] interface {
	Iterator
	Node() Node[K]
	Slice() []Node[K]
}

func NewNodes[K comparable](n map[K]Node[K]) Nodes[K] {
	ret := &nodes[K]{
		view: make([]Node[K], 0, len(n)),
		idx:  -1,
	}

	for _, v := range n {
		ret.view = append(ret.view, v)
	}

	return ret
}

type nodes[K comparable] struct {
	view []Node[K]
	idx  int
}

// Next returns whether the next call of Node will return a valid node.
func (n *nodes[K]) Next() bool {
	if n.idx+1 >= len(n.view) {
		n.idx = len(n.view)
		return false
	}

	n.idx++
	return true
}

// Len returns the remaining number of nodes to be iterated over.
func (n nodes[K]) Len() int {
	return len(n.view) - n.idx - 1
}

func (n *nodes[K]) Reset() {
	n.idx = -1
}

func (n nodes[K]) Node() Node[K] {
	if n.idx >= len(n.view) || n.idx < 0 {
		return nil
	}
	return n.view[n.idx]
}

func (n *nodes[K]) Slice() []Node[K] {
	if n.idx >= len(n.view) {
		return nil
	}
	idx := n.idx + 1
	n.idx = len(n.view)
	return n.view[idx:]
}
