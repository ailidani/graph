package graph

import (
	stdMath "math"

	"github.com/ailidani/graph/math"
	"github.com/ailidani/graph/queue"
)

/*
*  0 -> 1 -> 2 -> 3 -> 4
 */
func Line(n int) Graph[int] {
	g := New[int]()
	for i := 0; i < n; i++ {
		g.Add(i)
	}
	for i := 0; i < n-1; i++ {
		g.Connect(i, i+1)
	}
	return g
}

/*
*  0 -> 1 -> 2 -> 3 -> 4
*  |_ _ _ _ _ _ _ _ _ _|
 */
func Ring(n int) Graph[int] {
	g := New[int]()
	for i := 0; i < n; i++ {
		g.Add(i)
	}
	for i := 0; i < n; i++ {
		g.Connect(i, (i+1)%n)
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
func Star(n int) Graph[int] {
	g := New[int]()
	for i := 0; i < n; i++ {
		g.Add(i)
	}
	for i := 1; i < n; i++ {
		g.Connect(0, i)
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
func Tree(n int) Graph[int] {
	g := New[int]()
	for i := 0; i < n; i++ {
		g.Add(i)
	}
	for i := 1; i < n; i++ {
		g.Connect((i-1)/2, i)
	}
	return g
}

/*
* 0 - 1
* | X |
* 3 - 2
 */
// Kn complete graph
func K(n int) Graph[int] {
	g := New[int]()
	for i := 0; i < n; i++ {
		g.Add(i)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			g.Connect(i, j)
			g.Connect(j, i)
		}
	}
	return g
}

/*
*   4 - 5
*  /   /|
* 0 - 1 |
* |   |/
* 3 - 2
 */
// Qn hypercube graph
func Q(n int) Graph[int] {
	g := New[int]()
	for i := 0; i < 1<<uint(n); i++ {
		g.Add(i)
	}
	for i := 0; i < 1<<uint(n); i++ {
		for j := 0; j < n; j++ {
			g.Connect(i, i^(1<<uint(j)))
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
	for i := 0; i < 6; i++ {
		g.Add(i)
	}
	g.Connect(0, 1)
	g.Connect(0, 4)
	g.Connect(1, 2)
	g.Connect(1, 4)
	g.Connect(2, 3)
	g.Connect(3, 4)
	g.Connect(3, 5)
	return g
}

// PrimeSumGraph generates a graph where two integers in range [from, to] has an edge if their sum is a prime number
func PrimeSumGraph(from, to int) Graph[int] {
	g := New[int]()
	for i := from; i <= to; i++ {
		g.Add(i)
	}
	for i := from; i <= to; i++ {
		for j := i + 1; j <= to; j++ {
			if math.IsPrime(i + j) {
				g.Connect(i, j)
				g.Connect(j, i)
			}
		}
	}
	return g
}

// PrimeFactorGraph generates a graph where each integer in range has incoming edge from all its prime factors
func PrimeFactorGraph(from, to int) Graph[int] {
	g := New[int]()
	for i := from; i <= to; i++ {
		for _, p := range math.PrimeFactors(i) {
			// skip self edge
			if p == i {
				continue
			}
			if g.Edge(p, i) == nil {
				g.AddEdge(NewEdge(NewNode(p), NewNode(i)))
			}
		}
	}
	return g
}

// GoldbachGraph generates a graph where each even number has 2 incoming edges from 2 primes that sum equal to it
func GoldbachGraph(from, to int) Graph[int] {
	g := New[int]()
	for i := from; i <= to; i++ {
		for j := from; j <= i; j++ {
			if math.IsPrime(i) && math.IsPrime(j) && (i+j)%2 == 0 {
				g.AddEdge(NewEdge(NewNode(i), NewNode(i+j)))
				if i != j {
					g.AddEdge(NewEdge(NewNode(j), NewNode(i+j)))
				}
			}
		}

	}
	return g
}

// CollatzGraph...
func CollatzGraph(k int) Graph[int] {
	g := New[int]()
	queue := queue.New[int]()
	queue.Push(1)

	for queue.Size() > 0 && k > 0 {
		k--
		n := queue.Pop()
		if n%6 == 4 {
			i := (n - 1) / 3
			queue.Push(i)
			if g.Edge(i, n) == nil {
				g.AddEdge(NewEdge(NewNode(i), NewNode(n)))
			}
		}

		if n >= stdMath.MaxInt64/2 {
			// skip 2*n that overflow
			continue
		}
		j := 2 * n
		queue.Push(j)
		if g.Edge(j, n) == nil {
			g.AddEdge(NewEdge(NewNode(j), NewNode(n)))
		}
	}
	return g
}
