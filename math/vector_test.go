package math

import (
	"testing"
)

func TestMinMax(t *testing.T) {
	s := []float64{5.0, 10.0, 0.0}

	if Min(s) != 0.0 {
		t.Errorf("Min(%v) != 0.0", s)
	}

	if Max(s) != 10.0 {
		t.Errorf("Max(%v) != 10.0", s)
	}

	if MaxIndex(s) != 1 {
		t.Errorf("MaxIndex(%v) != 1", s)
	}
}

func TestDot(t *testing.T) {
	if Dot([]float64{1.0, 6.0, 3.0}, []float64{2.0, 2.0, 1.0}) != 17.0 {
		t.Errorf("Dot([]float64{1.0, 6.0, 3.0}, []float64{2.0, 2.0, 1.0}) != 17.0")
	}
}
