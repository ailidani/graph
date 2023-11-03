package graph

import (
	"slices"
)

func Degree[K comparable](g Graph[K], v K) int {
	return OutDegree(g, v) + InDegree(g, v)
}

func OutDegree[K comparable](g Graph[K], v K) int {
	return len(g.FromMap(v))
}

func InDegree[K comparable](g Graph[K], v K) int {
	return len(g.ToMap(v))
}

func DegreeSequence[K comparable](g Graph[K]) []int {
	degrees := make([]int, 0, g.Order())
	nodes := g.Nodes()
	for nodes.Next() {
		degrees = append(degrees, Degree(g, nodes.Node().ID()))
	}
	// slices.Sort(degrees)
	// slices.Reverse(degrees)
	// sort in descending order
	slices.SortFunc(degrees, func(a, b int) int { return b - a })
	return degrees
}
