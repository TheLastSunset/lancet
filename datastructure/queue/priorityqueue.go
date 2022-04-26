package datastructure

import "github.com/duke-git/lancet/v2/lancetconstraints"

// PriorityQueue is a implemented by binary heap tree
// which are referred to as the left child and the right child.
// type T should implements Compare function in lancetconstraints.Comparator interface.
type PriorityQueue[T any] struct {
	items      []T
	size       int
	comparator lancetconstraints.Comparator
}

// NewPriorityQueue return a pointer of PriorityQueue
func NewPriorityQueue[T any](capacity int, comparator lancetconstraints.Comparator) *NewPriorityQueue[T] {
	return &NewPriorityQueue[T]{
		items:      make([]T, capacity+1),
		size:       0,
		comparator: comparator,
	}
}