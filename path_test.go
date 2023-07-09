package graph

import (
	"testing"
)

func TestPathsBetween(t *testing.T) {
	t.Parallel()

	tests := [...]struct {
		name string
		Graph[int]
		dst   int
		count int
	}{
		{"line", Line(5), 4, 1},
		{"Q2", Q(2), 3, 2},
		{"K4", K(4), 3, 5},
		{"loop", Cycle(5), 4, 1},
		{"tree", Tree(5), 4, 1},
		{"star", Star(5), 4, 1},
	}

	for _, test := range tests {
		paths := PathsBetween(test.Graph, test.Graph.Node(0), test.Graph.Node(test.dst))
		if len(paths) != test.count {
			t.Errorf("test %s = %v", test.name, paths)
		}
	}
}

func TestPaths(t *testing.T) {
	t.Parallel()

	tests := [...]struct {
		string
		Graph[int]
		int
	}{
		{"line", Line(5), 1},
		{"loop", Cycle(5), 1},
		{"Q2", Q(2), 6},
		{"K4", K(4), 15},
		{"tree", Tree(5), 3},
		{"star", Star(5), 4},
	}

	for _, test := range tests {
		paths := Paths(test.Graph, test.Graph.Node(0))
		if len(paths) != test.int {
			t.Errorf("test %s = %v", test.string, paths)
		}
	}
}
