package graph

import "testing"

/*
 * 0 -> 1 -> 3
 * |    ^    |
 * |    |    |
 * +--> 2 <--+
 */

func TestGraph(t *testing.T) {
	g := New[int]()
	g.AddNode(NewNode(0))
	g.AddNode(NewNode(1))
	g.AddNode(NewNode(2))
	g.AddNode(NewNode(3))
	g.AddEdge(NewEdge(g.Node(0), g.Node(1)))
	g.AddEdge(NewEdge(g.Node(0), g.Node(2)))
	g.AddEdge(NewEdge(g.Node(1), g.Node(3)))
	g.AddEdge(NewEdge(g.Node(2), g.Node(1)))
	g.AddEdge(NewEdge(g.Node(3), g.Node(2)))

	if g.Order() != 4 {
		t.Errorf("Order() = %d, want %d", g.Order(), 4)
	}

	if g.Size() != 5 {
		t.Errorf("Size() = %d, want %d", g.Size(), 5)
	}

	if len(g.From(0)) != 2 {
		t.Errorf("From(0) = %d, want %d", len(g.From(0)), 2)
	}

	if len(g.To(0)) != 0 {
		t.Errorf("To(0) = %d, want %d", len(g.To(0)), 0)
	}

	bfs := []int{0, 1, 2, 3}
	result := make([]int, 0)
	BFS(g, g.Node(0), func(n Node[int], depth int) {
		result = append(result, n.ID())
	})
	for i, v := range result {
		if int(v) != bfs[i] {
			t.Fatalf("graph BFS(0) = %v", result)
		}
	}
}
