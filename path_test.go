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
		{"line", line(5), 4, 1},
		{"Q2", Q(2), 3, 2},
		{"K4", K(4), 3, 5},
		{"loop", loop(5), 4, 1},
		{"tree", tree(5), 4, 1},
		{"star", star(5), 4, 1},
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
		{"line", line(5), 1},
		{"loop", loop(5), 1},
		{"Q2", Q(2), 6},
		{"K4", K(4), 15},
		{"tree", tree(5), 3},
		{"star", star(5), 5},
	}

	for _, test := range tests {
		paths := Paths(test.Graph, test.Graph.Node(0))
		if len(paths) != test.int {
			t.Errorf("test %s = %v", test.string, paths)
		}
	}
}
