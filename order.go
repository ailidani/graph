package graph

// Reverse reverses the order of nodes.
func Reverse[K comparable](nodes []Node[K]) {
	for i, j := 0, len(nodes)-1; i < j; i, j = i+1, j-1 {
		nodes[i], nodes[j] = nodes[j], nodes[i]
	}
}
