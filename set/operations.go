package set

import "math"

func Equal[T comparable](a, b Set[T]) bool {
	if a.Size() != b.Size() {
		return false
	}

	switch v := a.(type) {
	case set[T]:
		for e := range v {
			if !b.Contains(e) {
				return false
			}
		}
		return true
	case MultiSet[T]:
		for e, c := range v {
			if b.(MultiSet[T])[e] != c {
				return false
			}
		}
		return true
	default:
		panic("unknown set type")
	}
}

func Merge[T comparable](a, b Set[T]) Set[T] {
	c := New[T]()
	for _, e := range a.Array() {
		c.Add(e)
	}
	for _, e := range b.Array() {
		c.Add(e)
	}
	return c
}

func PowerSet[T comparable](s Set[T]) []Set[T] {
	// total number of subsets
	n := int(math.Pow(2, float64(s.Size())))
	base := s.Array()
	pset := make([]Set[T], 0)

	for i := 0; i < n; i++ {
		subset := New[T]()
		// check every bit of i
		for j := 0; j < n; j++ {
			// if jth bit of i is set add s[j]
			if (i & (1 << j)) != 0 {
				subset.Add(base[j])
			}
		}
		pset = append(pset, subset)
	}

	return pset
}
