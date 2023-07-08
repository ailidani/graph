package graph

func Connected[K comparable](g Graph[K]) bool {
	visited := make(map[K]bool)
	nodes := g.Nodes()
	for nodes.Next() {
		n := nodes.Node()
		if visited[n.ID()] {
			continue
		}
		visited[n.ID()] = true
		BFS(g, n, func(n Node[K], depth int) {
			visited[n.ID()] = true
		})
	}
	nodes.Reset()
	return len(visited) == nodes.Len()
}

func Reachable[K comparable](g Graph[K], from, to Node[K]) bool {
	v := DFSUntil(g, from, func(n Node[K]) bool {
		return n == to
	})
	return v != nil
}
