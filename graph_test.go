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

/*
*  0 -> 1 -> 2 -> 3 -> 4
 */
func line(n int) Graph[int] {
	g := New[int]()
	for i := 0; i < n; i++ {
		g.AddNode(NewNode(i))
	}
	for i := 0; i < n-1; i++ {
		g.AddEdge(NewEdge(g.Node(i), g.Node(i+1)))
	}
	return g
}

/*
*  0 -> 1 -> 2 -> 3 -> 4
*  |_ _ _ _ _ _ _ _ _ _|
 */
func loop(n int) Graph[int] {
	g := New[int]()
	for i := 0; i < n; i++ {
		g.AddNode(NewNode(i))
	}
	for i := 0; i < n; i++ {
		g.AddEdge(NewEdge(g.Node(i), g.Node((i+1)%n)))
	}
	return g
}

/*
*      2
*      |
*  1 - 0 - 3
*     / \
*    5   4
 */
func star(n int) Graph[int] {
	g := New[int]()
	for i := 0; i < n; i++ {
		g.AddNode(NewNode(i))
	}
	for i := 1; i < n; i++ {
		g.AddEdge(NewEdge(g.Node(0), g.Node(i)))
	}
	return g
}

/*
*     0
*    / \
*   1   2
*  / \   \
* 3   4   5
 */
func tree(n int) Graph[int] {
	g := New[int]()
	for i := 0; i < n; i++ {
		g.AddNode(NewNode(i))
	}
	for i := 1; i < n; i++ {
		g.AddEdge(NewEdge(g.Node((i-1)/2), g.Node(i)))
	}
	return g
}

/*
0 - 1
| X |
3 - 2
*/
// Kn complete graph
func K(n int) Graph[int] {
	g := New[int]()
	for i := 0; i < n; i++ {
		g.AddNode(NewNode(i))
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			g.AddEdge(NewEdge(g.Node(i), g.Node(j)))
		}
	}
	return g
}

// Qn hypercube graph
func Q(n int) Graph[int] {
	g := New[int]()
	for i := 0; i < 1<<uint(n); i++ {
		g.AddNode(NewNode(i))
	}
	for i := 0; i < 1<<uint(n); i++ {
		for j := 0; j < n; j++ {
			g.AddEdge(NewEdge(g.Node(i), g.Node(i^(1<<uint(j)))))
		}
	}
	return g
}

/*
 *   0
 *  / \
 * 1 - 4
 * |   |
 * 2 - 3 - 5
 */
func BronKerboschGraph() Graph[int] {
	g := New[int]()
	g.AddNode(NewNode(0))
	g.AddNode(NewNode(1))
	g.AddNode(NewNode(2))
	g.AddNode(NewNode(3))
	g.AddNode(NewNode(4))
	g.AddNode(NewNode(5))

	g.AddEdge(NewEdge(g.Node(0), g.Node(1)))
	g.AddEdge(NewEdge(g.Node(0), g.Node(4)))
	g.AddEdge(NewEdge(g.Node(1), g.Node(2)))
	g.AddEdge(NewEdge(g.Node(1), g.Node(4)))
	g.AddEdge(NewEdge(g.Node(2), g.Node(3)))
	g.AddEdge(NewEdge(g.Node(3), g.Node(4)))
	g.AddEdge(NewEdge(g.Node(3), g.Node(5)))
	return g
}
