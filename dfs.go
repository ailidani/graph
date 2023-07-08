package graph

import (
	"github.com/ailidani/graph/set"
	"github.com/ailidani/graph/stack"
)

//---------------------------
// DFS
//---------------------------

// DFS visits depth first search nodes from a given source
func DFS[K comparable](g Graph[K], root Node[K], visit func(Node[K])) {
	if g.Node(root.ID()) == nil {
		panic("root not found")
	}

	visited := set.New[K]()
	stack := stack.New[Node[K]]()
	stack.Push(root)

	for !stack.Empty() {
		u := stack.Pop()
		if !visited.Contains(u.ID()) {
			visited.Add(u.ID())
			if visit != nil {
				visit(u)
			}
		}

		for _, v := range g.From(u.ID()) {
			if !visited.Contains(v.ID()) {
				stack.Push(v)
			}
		}
	}
}

// DFSUntil visits depth first search nodes from a given source, and stops if visitUntil function returns true
func DFSUntil[K comparable](g Graph[K], root Node[K], visitUntil func(Node[K]) bool) Node[K] {
	if g.Node(root.ID()) == nil {
		panic("root not found")
	}

	visited := set.New[K]()
	stack := stack.New[Node[K]]()
	stack.Push(root)

	for !stack.Empty() {
		u := stack.Pop()
		if !visited.Contains(u.ID()) {
			visited.Add(u.ID())
			if visitUntil != nil && visitUntil(u) {
				return u
			}
		}

		for _, v := range g.From(u.ID()) {
			if !visited.Contains(v.ID()) {
				stack.Push(v)
			}
		}
	}
	return nil
}
