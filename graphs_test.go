package graph

import (
	"os"
	"testing"
)

func TestCollatzGraph(t *testing.T) {
	g := CollatzGraph(1000)
	err := os.WriteFile("collatz.dot", DOT(g), 0644)
	if err != nil {
		panic(err)
	}
}

func TestPrimeSumGraph(t *testing.T) {
	g := PrimeSumGraph(1, 100)
	err := os.WriteFile("prime_sum.dot", DOT(g), 0644)
	if err != nil {
		panic(err)
	}
}
