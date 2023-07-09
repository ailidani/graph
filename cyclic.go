package graph

//---------------------------
// Cyclic
//---------------------------

type color int

const (
	white color = iota // unvisited nodes
	gray               // visiting nodes
	black              // visited nodes
)

func visit[K comparable](g Graph[K], v Node[K], colors map[K]color) bool {
	colors[v.ID()] = gray
	for _, u := range g.From(v.ID()) {
		if colors[u.ID()] == gray {
			return true
		}
		if colors[u.ID()] == white && visit(g, u, colors) {
			return true
		}
	}
	colors[v.ID()] = black
	return false
}

// Cyclic returns true if the graph contains a cycle
func Cyclic[K comparable](g Graph[K]) bool {
	colors := make(map[K]color)
	// set all nodes color to white
	nodes := g.Nodes()
	for nodes.Next() {
		v := nodes.Node()
		colors[v.ID()] = white
	}

	nodes.Reset()
	for nodes.Next() {
		v := nodes.Node()
		if colors[v.ID()] == white {
			if visit(g, v, colors) {
				return true
			}
		}
	}
	return false
}

func CyclicNode[K comparable](g Graph[K], v Node[K]) bool {
	colors := make(map[K]color)
	// set all nodes color to white
	nodes := g.Nodes()
	for nodes.Next() {
		u := nodes.Node()
		colors[u.ID()] = white
	}
	return visit(g, v, colors)
}

// CyclePath returns the first cycle with vertices
func CyclePath[K comparable](g Graph[K]) []Node[K] {
	colors := make(map[K]color)
	// set all nodes color to white
	nodes := g.Nodes()
	for nodes.Next() {
		v := nodes.Node()
		colors[v.ID()] = white
	}

	nodes.Reset()
	for nodes.Next() {
		v := nodes.Node()
		if colors[v.ID()] == white {
			if visit(g, v, colors) {
				cycle := make([]Node[K], 0)
				for u, color := range colors {
					if color == gray {
						cycle = append(cycle, g.Node(u))
					}
				}
				return cycle
			}
		}
	}
	return nil
}
