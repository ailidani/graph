package graph

import (
	"reflect"
	"testing"
)

/*
 * 0 -> 1 -> 3
 * |    ^    |
 * |    |    |
 * +--> 2 <--+
 */

func TestGraph(t *testing.T) {
	g := New[int]()
	g.Add(0)
	g.Add(1)
	g.Add(2)
	g.Add(3)
	g.Connect(0, 1)
	g.Connect(0, 2)
	g.Connect(1, 3)

	if g.Order() != 4 {
		t.Errorf("Order() = %d, want %d", g.Order(), 4)
	}

	if g.Size() != 3 {
		t.Errorf("Size() = %d, want %d", g.Size(), 3)
	}

	if len(g.From(0)) != 2 {
		t.Errorf("From(0) = %d, want %d", len(g.From(0)), 2)
	}

	if len(g.To(0)) != 0 {
		t.Errorf("To(0) = %d, want %d", len(g.To(0)), 0)
	}

	bfs := []int{0, 1, 2, 3}
	result := make([]int, 0)
	BFS(g, 0, func(n Node[int], d int) {
		result = append(result, n.ID())
	})
	for i, v := range result {
		if int(v) != bfs[i] {
			t.Fatalf("graph BFS(0) = %v", result)
		}
	}

	g.Connect(2, 1)
	g.Connect(3, 2)
	if !Cyclic(g) {
		t.Error("graph cannot detect cycle")
	}

	if !CyclicNode(g, 0) {
		t.Error("0 should not be part of cycle")
	}

	if !CyclicNode(g, 1) {
		t.Error("1 should be part of cycle")
	}

	pathes := PathsBetween(g, g.Node(0), g.Node(3))
	if len(pathes) != 2 {
		t.Error("0 to 3 should have 2 path")
	}
	for _, path := range pathes {
		result := make([]int, len(path))
		for i, n := range path {
			result[i] = n.ID()
		}
		if !reflect.DeepEqual(result, []int{0, 1, 3}) && !reflect.DeepEqual(result, []int{0, 2, 1, 3}) {
			t.Errorf("wrong path %v", path)
		}
	}

	pathes = PathsBetween(g, g.Node(0), g.Node(1))
	if len(pathes) != 2 {
		t.Error("0 to 1 should have 2 path")
	}

	pathes = PathsBetween(g, g.Node(0), g.Node(2))
	if len(pathes) != 2 {
		t.Error("0 to 2 should have 2 path")
	}
}

/*
 * 0 --> 1 <--> 4
 * ^     |
 * |     |
 * V     V
 * 3 <-- 2 <--> 5
 */
func TestCycle(t *testing.T) {
	g := New[int]()
	g.Add(0)
	g.AddEdge(NewEdge(NewNode(0), NewNode(1)))
	g.AddEdge(NewEdge(NewNode(1), NewNode(2)))
	g.AddEdge(NewEdge(NewNode(1), NewNode(4)))
	g.AddEdge(NewEdge(NewNode(4), NewNode(1)))
	g.AddEdge(NewEdge(NewNode(2), NewNode(3)))
	g.AddEdge(NewEdge(NewNode(3), NewNode(0)))
	g.AddEdge(NewEdge(NewNode(0), NewNode(3)))
	g.AddEdge(NewEdge(NewNode(2), NewNode(5)))
	g.AddEdge(NewEdge(NewNode(5), NewNode(2)))

	pathes := PathsBetween(g, g.Node(0), g.Node(3))
	if len(pathes) != 5 {
		t.Errorf("0 to 3 should have 2 path, got %d", len(pathes))
		t.Errorf("pathes = %v", pathes)
	}
}
