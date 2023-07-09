package graph

import (
	"math"
	"testing"
)

func TestDiameter(t *testing.T) {
	t.Parallel()

	tests := [...]struct {
		string
		Graph[int]
		int
	}{
		{"line", Line(5), int(math.Inf(1))},
		{"Q2", Q(2), 2},
		{"K4", K(4), 1},
		{"loop", Cycle(5), 4},
		{"tree", Tree(5), int(math.Inf(1))},
		{"star", Star(5), int(math.Inf(1))},
	}

	for _, test := range tests {
		d := Diameter(test.Graph)
		if int(d) != test.int {
			t.Errorf("test %s Diameter() = %v, want %v", test.string, d, test.int)
		}
	}
}
