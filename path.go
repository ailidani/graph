package graph

import (
	"slices"

	"github.com/ailidani/graph/queue"
)

//---------------------------
// Path between
//---------------------------

// PathsBetween finding all paths from source to destination
func PathsBetween[K comparable](g Graph[K], s, d Node[K]) [][]Node[K] {
	all := make([][]Node[K], 0)
	queue := queue.New[[]Node[K]]()
	// current path
	path := make([]Node[K], 0)
	path = append(path, s)
	queue.Push(path)
	for queue.Size() != 0 {
		path := queue.Pop()
		last := path[len(path)-1]
		if last.ID() == d.ID() {
			all = append(all, path)
			continue
		}
		for _, v := range g.From(last.ID()) {
			if slices.Contains(path, v) {
				continue
			}
			newPath := make([]Node[K], len(path))
			copy(newPath, path)
			newPath = append(newPath, v)
			queue.Push(newPath)
		}
	}

	return all
}

//---------------------------
// Path all
//---------------------------

func Paths[K comparable](g Graph[K], root Node[K]) [][]Node[K] {
	all := make([][]Node[K], 0)
	queue := queue.New[[]Node[K]]()
	// current path
	path := make([]Node[K], 0)
	path = append(path, root)
	queue.Push(path)
	for queue.Size() != 0 {
		path := queue.Pop()
		last := path[len(path)-1]
		// leaf node
		if len(g.From(last.ID())) == 0 {
			all = append(all, path)
			continue
		}
		for _, v := range g.From(last.ID()) {
			// cycle path
			if slices.Contains(path, v) {
				if !containsPath(all, path) {
					all = append(all, path)
				}
				continue
			}
			newPath := make([]Node[K], len(path))
			copy(newPath, path)
			newPath = append(newPath, v)
			queue.Push(newPath)
		}
	}

	return all
}

func containsPath[K comparable](paths [][]Node[K], path []Node[K]) bool {
	for _, i := range paths {
		if slices.EqualFunc(i, path, func(a, b Node[K]) bool {
			return a.ID() == b.ID()
		}) {
			return true
		}
	}
	return false
}
