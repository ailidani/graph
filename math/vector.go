package math

import (
	"errors"

	"golang.org/x/exp/constraints"
)

func Sum[T constraints.Ordered](x []T) (sum T) {
	for _, v := range x {
		sum += v
	}
	return
}

func Min[T constraints.Ordered](x []T) T {
	min := x[0]
	for _, v := range x {
		if v < min {
			min = v
		}
	}
	return min
}

func Max[T constraints.Ordered](x []T) T {
	max := x[0]
	for _, v := range x {
		if v > max {
			max = v
		}
	}
	return max
}

// MaxIndex returns the index of the largest element
func MaxIndex[T constraints.Ordered](x []T) int {
	max, idx := x[0], 0
	for i, v := range x {
		if v > max {
			max, idx = v, i
		}
	}
	return idx
}

func Dot[T constraints.Float](x, y []T) T {
	var p T
	for i := range x {
		p += x[i] * y[i]
	}
	return p
}

func MatrixToVector[T constraints.Float](x [][]T) ([]T, error) {
	if len(x) == 0 {
		return nil, errors.New("not matrix")
	}

	m := len(x)
	n := len(x[0])
	size := m * n

	var vector []T
	for _, row := range x {
		if len(row) != n {
			return nil, errors.New("not matrix")
		}
		vector = append(vector, row...)
	}

	if len(vector) != size {
		return nil, errors.New("vector size")
	}

	return vector, nil
}
