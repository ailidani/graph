package graph

type Graph[K comparable] interface {
	// Order of a graph is the number of vertices in the graph
	Order() int
	// Size of a graph is the number of edges in the graph
	Size() int
	Node(K) Node[K]
	Edge(from K, to K) Edge[K]
	Nodes() Nodes[K]
	Edges() Edges[K]
	From(K) []Node[K]
	FromMap(K) map[K]Edge[K]
	To(K) []Node[K]
	ToMap(K) map[K]Edge[K]
	Builder[K]
}

type WeightedGraph[K comparable] interface {
	Graph[K]
	WeightedEdge(K, K) WeightedEdge[K]
	Weight(K, K) (float64, bool)
}

type Builder[K comparable] interface {
	// AddNode adds a node to the graph, panics if node ID exist.
	AddNode(Node[K]) error
	RemoveNode(K) error
	// AddEdge adds an edge to the graph, adds from and to nodes if not exist.
	AddEdge(Edge[K]) error
	RemoveEdge(from K, to K) error
}

type graph[K comparable] struct {
	nodes map[K]Node[K]
	from  map[K]map[K]Edge[K]
	to    map[K]map[K]Edge[K]
}

func New[K comparable]() Graph[K] {
	return &graph[K]{
		nodes: make(map[K]Node[K]),
		from:  make(map[K]map[K]Edge[K]),
		to:    make(map[K]map[K]Edge[K]),
	}
}

//---------------------------
// Graph
//---------------------------

func (g graph[K]) Order() int { return len(g.nodes) }

func (g graph[K]) Size() int {
	sum := 0
	for _, edges := range g.from {
		sum += len(edges)
	}
	return sum
}

func (g graph[K]) Node(id K) Node[K] {
	return g.nodes[id]
}

func (g graph[K]) Edge(from, to K) Edge[K] {
	edges, exist := g.from[from]
	if !exist {
		return nil
	}

	edge, exist := edges[to]
	if !exist {
		return nil
	}

	return edge
}

func (g graph[K]) Nodes() Nodes[K] {
	if len(g.nodes) == 0 {
		return nil
	}
	return NewNodes(g.nodes)
}

func (g graph[K]) Edges() Edges[K] {
	if len(g.nodes) == 0 {
		return nil
	}
	return NewEdges(g.from)
}

func (g graph[K]) From(id K) []Node[K] {
	from := g.from[id]
	if len(from) == 0 {
		return nil
	}
	nodes := make([]Node[K], 0, len(from))
	for _, edge := range from {
		nodes = append(nodes, edge.To())
	}
	return nodes
}

func (g graph[K]) To(id K) []Node[K] {
	to := g.to[id]
	if len(to) == 0 {
		return nil
	}
	nodes := make([]Node[K], 0, len(to))
	for _, edge := range to {
		nodes = append(nodes, edge.From())
	}
	return nodes
}

func (g graph[K]) FromMap(id K) map[K]Edge[K] {
	return g.from[id]
}

func (g graph[K]) ToMap(id K) map[K]Edge[K] {
	return g.to[id]
}

func (g graph[K]) HasNode(id K) bool {
	_, exist := g.nodes[id]
	return exist
}

func (g graph[K]) HasEdge(from, to K) bool {
	if !g.HasNode(from) || !g.HasNode(to) {
		return false
	}
	_, exist := g.from[from][to]
	return exist
}

//---------------------------
// Graph Builder
//---------------------------

func (g *graph[K]) AddNode(n Node[K]) error {
	id := n.ID()
	if g.HasNode(id) {
		return ErrNodeExist
	}
	g.nodes[id] = n
	g.from[id] = make(map[K]Edge[K])
	g.to[id] = make(map[K]Edge[K])
	return nil
}

func (g *graph[K]) RemoveNode(id K) error {
	if !g.HasNode(id) {
		return nil
	}
	delete(g.nodes, id)
	delete(g.from, id)
	delete(g.to, id)
	for _, from := range g.from {
		delete(from, id)
	}
	for _, to := range g.to {
		delete(to, id)
	}
	return nil
}

func (g *graph[K]) AddEdge(e Edge[K]) error {
	from := e.From().ID()
	to := e.To().ID()

	if from == to {
		return ErrSelfEdge
	}

	if g.HasEdge(from, to) {
		return ErrEdgeExist
	}

	if !g.HasNode(from) {
		g.AddNode(e.From())
	}

	if !g.HasNode(to) {
		g.AddNode(e.To())
	}

	g.from[from][to] = e
	g.to[to][from] = e

	return nil
}

func (g *graph[K]) RemoveEdge(from, to K) error {
	if from == to {
		return ErrSelfEdge
	}

	if !g.HasNode(from) || !g.HasNode(to) {
		return nil
	}

	if !g.HasEdge(from, to) {
		return nil
	}

	delete(g.from[from], to)
	delete(g.to[to], from)

	return nil
}
