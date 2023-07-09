package graph

import (
	"reflect"
	"sort"
	"testing"
)

func TestBFS(t *testing.T) {
	tests := []struct {
		name  string
		graph Graph[int]
		want  [][]int
	}{
		{
			name:  "line",
			graph: Line(5),
			want:  [][]int{{0}, {1}, {2}, {3}, {4}},
		},
		{
			name:  "loop",
			graph: Cycle(5),
			want:  [][]int{{0}, {1}, {2}, {3}, {4}},
		},
		{
			name:  "start",
			graph: Star(5),
			want:  [][]int{{0}, {1, 2, 3, 4}},
		},
		{
			name:  "tree",
			graph: Tree(6),
			want:  [][]int{{0}, {1, 2}, {3, 4, 5}},
		},
		{
			name:  "BronKerboschGraph",
			graph: BronKerboschGraph(),
			want:  [][]int{{0}, {1, 4}, {2}, {3}, {5}},
		},
	}

	for _, test := range tests {
		var got [][]int
		BFS(test.graph, test.graph.Node(0), func(n Node[int], depth int) {
			if depth >= len(got) {
				got = append(got, []int(nil))
			}
			got[depth] = append(got[depth], n.ID())
		})
		for _, l := range got {
			sort.Slice(l, func(i, j int) bool {
				return l[i] < l[j]
			})
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("unexpected BFS level structure for test %s:\ngot:  %v\nwant: %v", test.name, got, test.want)
		}
	}
}
