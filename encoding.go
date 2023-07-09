package graph

import "fmt"

func DOT[K comparable](g Graph[K]) []byte {
	var dot string
	dot += "digraph graph {\n"
	nodes := g.Nodes()
	for nodes.Next() {
		n := nodes.Node()
		dot += fmt.Sprintf("%v;\n", n.ID())
	}
	edges := g.Edges()
	for edges.Next() {
		e := edges.Edge()
		dot += fmt.Sprintf("%v -> %v [label=\"%v\"];\n", e.From().ID(), e.To().ID(), e)
	}
	dot += "}\n"
	return []byte(dot)
}
