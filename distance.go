package graph

import "math"

func Eccentricity[K comparable](g Graph[K], u Node[K]) float64 {
	shortest := DijkstraFrom(g, u)
	var max float64 = math.Inf(-1)
	nodes := g.Nodes()
	for nodes.Next() {
		dist := shortest.WeightTo(nodes.Node())
		if dist > max {
			max = dist
		}
	}
	return max
}

func Radius[K comparable](g Graph[K]) float64 {
	var min float64 = math.Inf(1)
	nodes := g.Nodes()
	for nodes.Next() {
		e := Eccentricity(g, nodes.Node())
		if e < min {
			min = e
		}
	}
	return min
}

func Diameter[K comparable](g Graph[K]) float64 {
	var max float64 = math.Inf(-1)
	nodes := g.Nodes()
	for nodes.Next() {
		e := Eccentricity(g, nodes.Node())
		if e > max {
			max = e
		}
	}
	return max
}
