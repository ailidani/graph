package graph

func Reverse[K comparable](g Graph[K]) Graph[K] {
	reversed := New[K]()
	edges := g.Edges()
	for edges.Next() {
		edge := edges.Edge()
		reversed.AddEdge(NewEdge(edge.To(), edge.From()))
	}
	return reversed
}
