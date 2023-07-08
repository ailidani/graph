package graph

import "github.com/ailidani/graph/heap"

// DijkstraFrom returns a shortest-path tree for a shortest path from u to all nodes in
// the graph g. If the graph does not implement Weighted, UniformCost is used.
// DijkstraFrom will panic if g has a u-reachable negative edge weight.
//
// If g is a graph.Graph, all nodes of the graph will be stored in the shortest-path
// tree, otherwise only nodes reachable from u will be stored.
//
// The time complexity of DijkstraFrom is O(|E|.log|V|).
func DijkstraFrom[K comparable](g Graph[K], u Node[K]) Shortest[K] {
	var path Shortest[K]
	if g.From(u.ID()) == nil {
		return Shortest[K]{from: u}
	}
	path = newShortest(u, g.Nodes().Slice())

	var weight Weighting[K]
	if wg, ok := g.(WeightedGraph[K]); ok {
		weight = wg.Weight
	} else {
		weight = UniformCost(g)
	}

	// Dijkstra's algorithm here is implemented essentially as
	// described in Function B.2 in figure 6 of UTCS Technical
	// Report TR-07-54.
	//
	// This implementation deviates from the report as follows:
	// - the value of path.dist for the start vertex u is initialized to 0;
	// - outdated elements from the priority queue (i.e. with respect to the dist value)
	//   are skipped.
	//
	// http://www.cs.utexas.edu/ftp/techreports/tr07-54.pdf
	Q := heap.New(func(i, j distanceNode[K]) bool {
		return i.dist < j.dist
	})
	Q.Push(distanceNode[K]{node: u, dist: 0})
	for Q.Size() > 0 {
		mid, ok := Q.Pop()
		if !ok {
			panic("dijkstra: unexpected empty priority queue")
		}
		k := path.index[mid.node.ID()]
		if mid.dist > path.dist[k] {
			continue
		}
		mnid := mid.node.ID()
		to := g.From(mnid)
		for _, v := range to {
			vid := v.ID()
			j, ok := path.index[vid]
			if !ok {
				j = path.add(v)
			}
			w, ok := weight(mnid, vid)
			if !ok {
				panic("dijkstra: unexpected invalid weight")
			}
			if w < 0 {
				panic("dijkstra: negative edge weight")
			}
			joint := path.dist[k] + w
			if joint < path.dist[j] {
				Q.Push(distanceNode[K]{node: v, dist: joint})
				path.set(j, joint, k)
			}
		}
	}

	return path
}

type distanceNode[K comparable] struct {
	node Node[K]
	dist float64
}
