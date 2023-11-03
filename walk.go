package graph

import "math/rand"

func RandomWalk[K comparable](g Graph[K], start Node[K], steps int) []Node[K] {
	walk := make([]Node[K], steps)
	walk[0] = start

	for i := 1; i < steps; i++ {
		nodes := g.From(walk[i-1].ID())
		if len(nodes) == 0 {
			break
		}

		walk[i] = nodes[rand.Intn(len(nodes))]
	}

	return walk
}
