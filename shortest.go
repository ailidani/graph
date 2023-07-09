package graph

import (
	"math"
)

//----------
// Shortest
//----------

type Shortest[K comparable] struct {
	from  Node[K]
	nodes []Node[K]
	index map[K]int
	dist  []float64
	next  []int
}

func newShortest[K comparable](from Node[K], nodes []Node[K]) Shortest[K] {
	index := make(map[K]int, len(nodes))
	for i, n := range nodes {
		index[n.ID()] = i
	}
	p := Shortest[K]{
		from:  from,
		nodes: nodes,
		index: index,
		dist:  make([]float64, len(nodes)),
		next:  make([]int, len(nodes)),
	}
	for i := range nodes {
		p.dist[i] = math.Inf(1)
		p.next[i] = -1
	}
	p.dist[index[from.ID()]] = 0
	return p
}

func (s Shortest[K]) From() Node[K] { return s.from }

func (s Shortest[K]) To(v Node[K]) ([]Node[K], float64) {
	to, exist := s.index[v.ID()]
	if !exist || math.IsInf(s.dist[to], 1) {
		return nil, math.Inf(1)
	}
	from := s.index[s.from.ID()]
	path := []Node[K]{s.nodes[to]}
	weight := math.Inf(1)
	n := len(s.nodes)
	for to != from {
		to = s.next[to]
		path = append(path, s.nodes[to])
		if n < 0 {
			panic("unexpected negative cycle")
		}
		n--
	}
	Reverse(path)
	return path, math.Min(weight, s.dist[s.index[v.ID()]])
}

func (s Shortest[K]) WeightTo(to Node[K]) float64 {
	index, exist := s.index[to.ID()]
	if !exist {
		return math.Inf(1)
	}
	return s.dist[index]
}

func (s *Shortest[K]) add(u Node[K]) int {
	if _, exist := s.index[u.ID()]; exist {
		panic("node exist")
	}
	idx := len(s.nodes)
	s.index[u.ID()] = idx
	s.nodes = append(s.nodes, u)
	s.dist = append(s.dist, math.Inf(1))
	s.next = append(s.next, -1)
	return idx
}

func (s *Shortest[K]) set(to int, weight float64, next int) {
	s.dist[to] = weight
	s.next[to] = next
}
