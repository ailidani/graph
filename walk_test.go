package graph

import "testing"

func TestRandomWalk(t *testing.T) {
	tests := []struct {
		Graph[int]
		step int
		walk []int
	}{
		{Line(10), 10, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{Ring(10), 10, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{Star(10), 10, nil},
	}

	for _, test := range tests {
		if test.walk == nil {
			continue
		}

		walk := RandomWalk(test.Graph, test.Graph.Node(0), test.step)
		if len(walk) != len(test.walk) {
			t.Errorf("Expected walk of length %d, got %d", len(test.walk), len(walk))
		}
		for i := 0; i < len(walk); i++ {
			if walk[i] != test.Graph.Node(test.walk[i]) {
				t.Errorf("Expected walk[%d] to be node %d, got %d", i, test.walk[i], walk[i])
			}
		}
	}
}
