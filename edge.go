package graph

type Edge[K comparable] interface {
	From() Node[K]
	To() Node[K]
}

type WeightedEdge[K comparable] interface {
	Edge[K]
	Weight() float64
}

func NewEdge[K comparable](from Node[K], to Node[K]) Edge[K] {
	return &edge[K]{
		from: from,
		to:   to,
	}
}

func NewWeightedEdge[K comparable](from Node[K], to Node[K], weight float64) Edge[K] {
	return &edge[K]{
		from:   from,
		to:     to,
		weight: weight,
	}
}

//--------------------
// edge
//--------------------

type edge[K comparable] struct {
	from   Node[K]
	to     Node[K]
	weight float64
}

func (e edge[K]) From() Node[K]   { return e.from }
func (e edge[K]) To() Node[K]     { return e.to }
func (e edge[K]) Weight() float64 { return e.weight }
func (e edge[K]) Reverse() Edge[K] {
	return &edge[K]{
		from:   e.to,
		to:     e.from,
		weight: e.weight,
	}
}

//--------------------
// edges
//--------------------

type Edges[K comparable] interface {
	Iterator
	Edge() Edge[K]
}

func NewEdges[K comparable](e map[K]map[K]Edge[K]) Edges[K] {
	edges := &edges[K]{
		edges: make([]Edge[K], 0),
		idx:   -1,
	}
	for u := range e {
		for _, edge := range e[u] {
			edges.edges = append(edges.edges, edge)
		}
	}
	return edges
}

type edges[K comparable] struct {
	edges []Edge[K]
	idx   int
}

func (e *edges[K]) Next() bool {
	if uint(e.idx)+1 >= uint(len(e.edges)) {
		e.idx = len(e.edges)
		return false
	}
	e.idx++
	return true
}

func (e *edges[K]) Len() int {
	if e.idx >= len(e.edges) {
		return 0
	}
	if e.idx <= 0 {
		return len(e.edges)
	}
	return len(e.edges[e.idx:])
}

func (e *edges[K]) Edge() Edge[K] {
	if e.idx >= len(e.edges) || e.idx < 0 {
		return nil
	}
	return e.edges[e.idx]
}

func (e *edges[K]) Reset() {
	e.idx = -1
}
