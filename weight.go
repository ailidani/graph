package graph

import (
	"math"
)

type Weighting[K comparable] func(K, K) (float64, bool)

// UniformCost returns a Weighting that returns an edge cost of 1 for existing
// edges, zero for node identity and Inf for otherwise absent edges.
func UniformCost[K comparable](g Graph[K]) Weighting[K] {
	return func(from, to K) (float64, bool) {
		if from == to {
			return 0, true
		}
		if g.Edge(from, to) != nil {
			return 1, true
		}
		return math.Inf(1), false
	}
}
