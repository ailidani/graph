package graph

import (
	"testing"
)

func TestDijkstraFrom(t *testing.T) {
	t.Parallel()

	tests := [...]struct {
		string
		Graph[int]
	}{
		{"line", Line(5)},
	}

	for _, test := range tests {
		shortest := DijkstraFrom(test.Graph, test.Graph.Node(0))
		for _, node := range test.Graph.Nodes().Slice() {
			if shortest.WeightTo(node) != float64(node.ID()) {
				t.Errorf("test %s: shortest path to %d = %v", test.string, node.ID(), shortest.WeightTo(node))
			}
		}
	}
}
