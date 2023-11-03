package graph

import (
	"encoding/json"
	"fmt"
)

func Json[K comparable](g Graph[K]) []byte {
	type d3 struct {
		Nodes []struct{ ID K }
		Links []struct {
			Source K
			Target K
		}
	}

	d := d3{}
	nodes := g.Nodes()
	for nodes.Next() {
		n := nodes.Node()
		d.Nodes = append(d.Nodes, struct{ ID K }{n.ID()})
	}
	edges := g.Edges()
	for edges.Next() {
		e := edges.Edge()
		d.Links = append(d.Links, struct {
			Source K
			Target K
		}{e.From().ID(), e.To().ID()})
	}
	json, err := json.Marshal(d)
	//json, err := json.Marshal(g.(*graph).from)
	if err != nil {
		panic(err)
	}
	return json
}

func CSV[K comparable](g Graph[K]) []byte {
	var csv string
	// csv += "Id;Label\n"
	// nodes := g.Nodes()
	// for nodes.Next() {
	// 	n := nodes.Node()
	// 	csv += fmt.Sprintf("%v,%v\n", g.ID(n), n)
	// }
	csv += "Source;Target;Label\n"
	edges := g.Edges()
	for edges.Next() {
		e := edges.Edge()
		csv += fmt.Sprintf("%v;%v;%v\n", e.From().ID(), e.To().ID(), e)
	}
	return []byte(csv)
}

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
