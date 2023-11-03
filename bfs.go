package graph

import (
	"github.com/ailidani/graph/queue"
	"github.com/ailidani/graph/set"
)

//---------------------------
// BFS
//---------------------------

// BFS visits breath first search nodes from a given source node.
func BFS[K comparable](g Graph[K], root K, visit func(n Node[K], depth int)) {
	BFSUntil(g, root, func(n Node[K], depth int) bool {
		visit(n, depth)
		return false
	})
}

func BFSUntil[K comparable](g Graph[K], root K, visitUntil func(n Node[K], depth int) bool) Node[K] {
	if g.Node(root) == nil {
		panic("root not found")
	}

	visited := set.New[K]()
	queue := queue.New[K]()
	visited.Add(root)
	queue.Push(root)

	var depth int = 0
	var children int = 0
	var untilNext int = 1

	for queue.Size() > 0 {
		u := queue.Pop()
		if visitUntil != nil && visitUntil(g.Node(u), depth) {
			return g.Node(u)
		}

		// get all children
		for _, v := range g.From(u) {
			if !visited.Contains(v.ID()) {
				visited.Add(v.ID())
				queue.Push(v.ID())
				children++
			}
		}

		if untilNext--; untilNext == 0 {
			depth++
			untilNext = children
			children = 0
		}
	}
	return nil
}
