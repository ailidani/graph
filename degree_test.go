package graph

import (
	"slices"
	"testing"
)

func TestDegreeSequence(t *testing.T) {
	tree := Tree(10)
	degrees := DegreeSequence(tree)
	if len(degrees) != 10 {
		t.Errorf("Expected degree sequence to have length 10, got %d", len(degrees))
	}

	if !slices.IsSortedFunc(degrees, func(a, b int) int { return b - a }) {
		t.Errorf("Expected degree sequence to be sorted in descending order, got %v", degrees)
	}
}
