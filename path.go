package graph

import (
	"log"

	"github.com/ailidani/graph/list"
	"github.com/ailidani/graph/slice"
)

//---------------------------
// Path between
//---------------------------

// PathsBetween finding all paths from source to destination
func PathsBetween[K comparable](g Graph[K], s, d Node[K]) [][]Node[K] {
	all := make([][]Node[K], 0)
	path := new(list.List[Node[K]])
	visited := make(map[Node[K]]bool)
	path.PushBack(s)
	dfs(g, s, d, visited, path, &all)
	return all
}

func dfs[K comparable](g Graph[K], s, d Node[K], visited map[Node[K]]bool, path *list.List[Node[K]], all *[][]Node[K]) {
	if s == d {
		*all = append(*all, path.Front().ToArray())
		return
	}
	if path.Len() > g.Nodes().Len() {
		return
	}
	visited[s] = true
	for _, v := range g.From(s.ID()) {
		if visited[v] {
			continue
		}
		path.PushBack(v)
		dfs(g, v, d, visited, path, all)
		path.Remove(path.Back())
	}
	visited[s] = false
}

//---------------------------
// Path all
//---------------------------

func Paths[K comparable](g Graph[K], root Node[K]) [][]Node[K] {
	all := make([][]Node[K], 0)
	path := new(list.List[Node[K]])
	dfs2(g, root, path, &all)
	return all
}

func containsPath[K comparable](paths [][]Node[K], path []Node[K]) bool {
	for _, i := range paths {
		if slice.Equal(i, path, func(a, b Node[K]) bool {
			return a.ID() == b.ID()
		}) {
			return true
		}
	}
	return false
}

func dfs2[K comparable](g Graph[K], s Node[K], path *list.List[Node[K]], all *[][]Node[K]) {
	if s == nil {
		return
	}

	// unique node in path
	if path.Contains(s, func(a, b Node[K]) bool {
		return a.ID() == b.ID()
	}) {
		p := path.Front().ToArray()
		if !containsPath(*all, p) {
			*all = append(*all, p)
		}
		return
	}

	path.PushBack(s)

	// leaf node
	if len(g.From(s.ID())) == 0 {
		p := path.Front().ToArray()
		if !containsPath(*all, p) {
			*all = append(*all, p)
		}
		path.Remove(path.Back())
		return
	}

	if path.Len() > g.Order() {
		log.Printf("%v", path.Front().ToArray())
		log.Fatalf("path.Len()=%d, graph.NodeSize()=%d", path.Len(), g.Order())
	}

	for _, v := range g.From(s.ID()) {
		dfs2(g, v, path, all)
	}

	path.Remove(path.Back())
}
