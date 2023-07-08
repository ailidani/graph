package graph

import (
	"errors"
)

var (
	ErrNodeNotFound = errors.New("node not found")
	ErrEdgeNotFound = errors.New("edge not found")
	ErrNodeExist    = errors.New("node exist")
	ErrEdgeExist    = errors.New("edge exist")
	ErrSelfEdge     = errors.New("self edge")
)
