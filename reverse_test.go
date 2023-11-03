package graph

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name  string
		graph Graph[int]
	}{
		{"line", Line(5)},
	}

	for _, test := range tests {
		reversed := Reverse(test.graph)
		reversed2 := Reverse(reversed)
		// TODO need graph comparison function
		if !reflect.DeepEqual(test.graph, reversed2) {
			t.Errorf("expect %s, got %s", Json(test.graph), Json(reversed2))
		}
	}
}
