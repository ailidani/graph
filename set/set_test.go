package set

import (
	"math"
	"reflect"
	"testing"

	"github.com/ailidani/graph/slice"
)

func TestSet(t *testing.T) {
	s := New[int]()

	for i := 1; i <= 10; i++ {
		s.Add(i)
	}

	if !s.Contains(5) {
		t.Error("missing element")
	}

	s.Remove(5)
	if s.Contains(5) {
		t.Error("cannot remove element")
	}
}

func TestPowerSet(t *testing.T) {
	n := 3
	s := New[int]()

	for i := 1; i <= n; i++ {
		s.Add(i)
	}

	ps := PowerSet(s)
	expected := int(math.Pow(2, float64(n)))
	if len(ps) != expected {
		t.Errorf("wrong power set size expected: %d actual: %d", expected, len(ps))
	}
	if !slice.Contains(ps, New[int](), func(a, b Set[int]) bool {
		return reflect.DeepEqual(a, b)
	}) {
		t.Error("does not contain empty set")
	}
	if !slice.Contains(ps, s, func(a, b Set[int]) bool {
		return reflect.DeepEqual(a, b)
	}) {
		t.Error("does not contain self")
	}
	for _, subset := range ps {
		t.Log(subset.Array())
	}
}
