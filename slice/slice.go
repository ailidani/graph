package slice

import (
	"math/rand"
)

// EqualsFn is a function that returns whether 'a' and 'b' are equal.
type EqualsFn[T any] func(a, b T) bool

// Shuffle the slice
func Shuffle[T any](a []T) {
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
}

// Random pick a random entry from slice
func Random[T any](s []T) T {
	return s[rand.Intn(len(s))]
}

// Remove removes an entry from slice
func Remove[T comparable](s []T, r T) []T {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

// Copy returns a copy of the slice
func Copy[T any](s []T) []T {
	d := make([]T, len(s))
	copy(d, s)
	return d
}

// IndexOf returns index position in slice from given entry
// If value is -1, the entry does not found.
func IndexOf[T comparable](slice []T, s T) int {
	for i, v := range slice {
		if v == s {
			return i
		}
	}
	return -1
}

// Include returns true or false if given entry is in slice
func Include[T comparable](slice []T, s T) bool {
	return IndexOf(slice, s) != -1
}

// UniqueAppend appends an entry if not exist in the slice
func UniqueAppend[T comparable](slice []T, s ...T) []T {
	for i := range s {
		if IndexOf(slice, s[i]) != -1 {
			continue
		}
		slice = append(slice, s[i])
	}
	return slice
}

func Prepend[T any](slice []T, e T) []T {
	slice = append(slice, e)
	copy(slice[1:], slice)
	slice[0] = e
	return slice
}

// Equal checks if the slices are equal
func Equal[T any](a, b []T, equals EqualsFn[T]) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !equals(a[i], b[i]) {
			return false
		}
	}
	return true
}

// Reverse reverses the slice
func Reverse[T any](slice []T) []T {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

func Contains[T any](slice []T, s T, equals EqualsFn[T]) bool {
	for _, v := range slice {
		if equals(v, s) {
			return true
		}
	}
	return false
}
